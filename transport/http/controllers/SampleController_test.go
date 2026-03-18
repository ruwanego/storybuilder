package controllers

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/boundary/repositories"
	"github.com/storybuilder/storybuilder/domain/entities"
)

// MockSampleRepository is a mock of the SampleRepositoryInterface.
type MockSampleRepository struct {
	mock.Mock
}

func (m *MockSampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Sample), args.Error(1)
}

func (m *MockSampleRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Sample), args.Error(1)
}

func (m *MockSampleRepository) Add(ctx context.Context, sample entities.Sample) error {
	args := m.Called(ctx, sample)
	return args.Error(0)
}

func (m *MockSampleRepository) Edit(ctx context.Context, sample entities.Sample) error {
	args := m.Called(ctx, sample)
	return args.Error(0)
}

func (m *MockSampleRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockValidatorAdapter is a mock for the ValidatorAdapterInterface.
type MockValidatorAdapter struct {
	mock.Mock
}

func (m *MockValidatorAdapter) Validate(data any) map[string]string {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(map[string]string)
}

func (m *MockValidatorAdapter) ValidateField(field any, rules string) map[string]string {
	args := m.Called(field, rules)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(map[string]string)
}

// newTestSampleController builds up a SampleController from mocked dependencies.
func newTestSampleController(mockRepo repositories.SampleRepositoryInterface, mockValidator adapters.ValidatorAdapterInterface) *SampleController {
	var logBuf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuf, nil))
	ctr := &container.Container{
		Adapters: container.Adapters{
			LogAdapter:       logger,
			ValidatorAdapter: mockValidator,
		},
		Repositories: container.Repositories{
			SampleRepository: mockRepo,
		},
	}
	return NewSampleController(ctr)
}

// ----- Get -----

func TestSampleController_Get_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	samples := []entities.Sample{{ID: 1, Name: "foo"}}
	mockRepo.On("Get", mock.Anything).Return(samples, nil)

	ctl := newTestSampleController(mockRepo, mockValidator)

	resp, err := ctl.Get(context.Background(), &struct{}{})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Body.Data, 1)
	assert.Equal(t, 1, resp.Body.Data[0].ID)

	mockRepo.AssertExpectations(t)
}

func TestSampleController_Get_Error(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockRepo.On("Get", mock.Anything).Return(nil, errors.New("db error"))

	ctl := newTestSampleController(mockRepo, mockValidator)

	resp, err := ctl.Get(context.Background(), &struct{}{})
	assert.Error(t, err)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
}

// ----- GetByID -----

func TestSampleController_GetByID_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockRepo.On("GetByID", mock.Anything, 1).Return(entities.Sample{ID: 1, Name: "foo"}, nil)

	ctl := newTestSampleController(mockRepo, mockValidator)

	input := &GetSampleByIDInput{ID: 1}
	resp, err := ctl.GetByID(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.Body.Data.ID)
	assert.Equal(t, "foo", resp.Body.Data.Name)

	mockRepo.AssertExpectations(t)
}

// Validation failure is now handled automatically by huma

// ----- Delete -----

func TestSampleController_Delete_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockRepo.On("Delete", mock.Anything, 5).Return(nil)

	ctl := newTestSampleController(mockRepo, mockValidator)

	input := &DeleteSampleInput{ID: 5}
	resp, err := ctl.Delete(context.Background(), input)

	assert.NoError(t, err)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
}

func TestSampleController_Delete_Error(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockRepo.On("Delete", mock.Anything, 5).Return(errors.New("delete failed"))

	ctl := newTestSampleController(mockRepo, mockValidator)

	input := &DeleteSampleInput{ID: 5}
	resp, err := ctl.Delete(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, resp)

	mockRepo.AssertExpectations(t)
}
