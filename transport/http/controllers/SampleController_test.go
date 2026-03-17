package controllers

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
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

// requestWithChiID creates an http.Request with a chi URL param "id" injected.
func requestWithChiID(method, path, id string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
}

// ----- Get -----

func TestSampleController_Get_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	samples := []entities.Sample{{ID: 1, Name: "foo"}}
	mockRepo.On("Get", mock.Anything).Return(samples, nil)

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.Get)

	req, _ := http.NewRequest("GET", "/samples", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestSampleController_Get_Error(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockRepo.On("Get", mock.Anything).Return(nil, errors.New("db error"))

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.Get)

	req, _ := http.NewRequest("GET", "/samples", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}

// ----- GetByID -----

func TestSampleController_GetByID_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockValidator.On("ValidateField", 1, "required,gt=0").Return(nil)
	mockRepo.On("GetByID", mock.Anything, 1).Return(entities.Sample{ID: 1, Name: "foo"}, nil)

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.GetByID)

	req := requestWithChiID("GET", "/samples/1", "1")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestSampleController_GetByID_ValidationFail(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockValidator.On("ValidateField", 0, "required,gt=0").Return(map[string]string{"id": "must be greater than 0"})

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.GetByID)

	req := requestWithChiID("GET", "/samples/abc", "abc")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Validation failure => 422 Unprocessable Entity
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

// ----- Delete -----

func TestSampleController_Delete_Success(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockValidator.On("ValidateField", 5, "required,gt=0").Return(nil)
	mockRepo.On("Delete", mock.Anything, 5).Return(nil)

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.Delete)

	req := requestWithChiID("DELETE", "/samples/5", "5")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestSampleController_Delete_Error(t *testing.T) {
	mockRepo := new(MockSampleRepository)
	mockValidator := new(MockValidatorAdapter)

	mockValidator.On("ValidateField", 5, "required,gt=0").Return(nil)
	mockRepo.On("Delete", mock.Anything, 5).Return(errors.New("delete failed"))

	ctl := newTestSampleController(mockRepo, mockValidator)
	handler := ctl.Wrap(ctl.Delete)

	req := requestWithChiID("DELETE", "/samples/5", "5")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockRepo.AssertExpectations(t)
}
