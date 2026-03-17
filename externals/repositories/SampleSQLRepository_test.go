package repositories

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/storybuilder/storybuilder/domain/entities"
)

type MockDBAdapter struct {
	mock.Mock
}

func (m *MockDBAdapter) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDBAdapter) Query(ctx context.Context, query string, parameters map[string]any) ([]map[string]any, error) {
	args := m.Called(ctx, query, parameters)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]any), args.Error(1)
}

func (m *MockDBAdapter) NewTransaction() (*sql.Tx, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*sql.Tx), args.Error(1)
}

func (m *MockDBAdapter) Destruct() {
	m.Called()
}

func TestNewSampleSQLRepository(t *testing.T) {
	mockDB := new(MockDBAdapter)
	
	// Test default table name
	repo := NewSampleSQLRepository(mockDB, "")
	assert.NotNil(t, repo)
	assert.Equal(t, "sample", repo.(*SampleSQLRepository).tableName)

	// Test custom table name
	repoCustom := NewSampleSQLRepository(mockDB, "test_sample")
	assert.NotNil(t, repoCustom)
	assert.Equal(t, "test_sample", repoCustom.(*SampleSQLRepository).tableName)
}

func TestSampleSQLRepository_Get(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	expectedRes := []map[string]any{
		{"id": int64(1), "name": []byte("test")},
	}
	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample", map[string]any{}).Return(expectedRes, nil)

	samples, err := repo.Get(context.Background())
	assert.NoError(t, err)
	assert.Len(t, samples, 1)
	assert.Equal(t, 1, samples[0].ID)
	assert.Equal(t, "test", samples[0].Name)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_Get_Error(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample", map[string]any{}).Return(nil, errors.New("db error"))

	samples, err := repo.Get(context.Background())
	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	assert.Nil(t, samples)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_GetByID(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	expectedRes := []map[string]any{
		{"id": int64(2), "name": []byte("test2")},
	}
	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample WHERE id=?id", map[string]any{"id": 2}).Return(expectedRes, nil)

	sample, err := repo.GetByID(context.Background(), 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, sample.ID)
	assert.Equal(t, "test2", sample.Name)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_GetByID_NotFound(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample WHERE id=?id", map[string]any{"id": 10}).Return([]map[string]any{}, nil)

	sample, err := repo.GetByID(context.Background(), 10)
	assert.NoError(t, err)
	assert.Equal(t, 0, sample.ID) // Empty struct
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_GetByID_Error(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample WHERE id=?id", map[string]any{"id": 3}).Return(nil, errors.New("db find err"))

	sample, err := repo.GetByID(context.Background(), 3)
	assert.Error(t, err)
	assert.Equal(t, "db find err", err.Error())
	assert.Equal(t, 0, sample.ID)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_Add(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	smpl := entities.Sample{Name: "new", Password: "pass"}
	mockDB.On("Query", mock.Anything, "INSERT INTO sample (name, password) VALUES(?name, ?password)", map[string]any{"name": "new", "password": "pass"}).Return(nil, nil)

	err := repo.Add(context.Background(), smpl)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_Edit(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	smpl := entities.Sample{ID: 5, Name: "edited", Password: "newpass"}
	mockDB.On("Query", mock.Anything, "UPDATE sample SET name=?name, password=?password WHERE id=?id", map[string]any{"id": 5, "name": "edited", "password": "newpass"}).Return(nil, nil)

	err := repo.Edit(context.Background(), smpl)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_Delete(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	mockDB.On("Query", mock.Anything, "DELETE FROM sample WHERE id=?id", map[string]any{"id": 8}).Return(nil, nil)

	err := repo.Delete(context.Background(), 8)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestSampleSQLRepository_MapResultPanicRecover(t *testing.T) {
	mockDB := new(MockDBAdapter)
	repo := NewSampleSQLRepository(mockDB, "sample")

	// Missing name field to trigger type assertion panic: row["name"].([]byte)
	badRes := []map[string]any{
		{"id": int64(1)}, 
	}
	mockDB.On("Query", mock.Anything, "SELECT id, name, password FROM sample", map[string]any{}).Return(badRes, nil)

	samples, err := repo.Get(context.Background())
	assert.Error(t, err) // the recover should natively return the panic casted to an error
	assert.Nil(t, samples)
	mockDB.AssertExpectations(t)
}
