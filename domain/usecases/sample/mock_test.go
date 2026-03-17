package sample

import (
	"context"

	"github.com/storybuilder/storybuilder/domain/entities"
)

type mockSampleRepository struct {
	mockGet     func(ctx context.Context) ([]entities.Sample, error)
	mockGetByID func(ctx context.Context, id int) (entities.Sample, error)
	mockAdd     func(ctx context.Context, sample entities.Sample) error
	mockEdit    func(ctx context.Context, sample entities.Sample) error
	mockDelete  func(ctx context.Context, id int) error
}

func (m *mockSampleRepository) Get(ctx context.Context) ([]entities.Sample, error) {
	if m.mockGet != nil {
		return m.mockGet(ctx)
	}
	return nil, nil
}

func (m *mockSampleRepository) GetByID(ctx context.Context, id int) (entities.Sample, error) {
	if m.mockGetByID != nil {
		return m.mockGetByID(ctx, id)
	}
	return entities.Sample{}, nil
}

func (m *mockSampleRepository) Add(ctx context.Context, sample entities.Sample) error {
	if m.mockAdd != nil {
		return m.mockAdd(ctx, sample)
	}
	return nil
}

func (m *mockSampleRepository) Edit(ctx context.Context, sample entities.Sample) error {
	if m.mockEdit != nil {
		return m.mockEdit(ctx, sample)
	}
	return nil
}

func (m *mockSampleRepository) Delete(ctx context.Context, id int) error {
	if m.mockDelete != nil {
		return m.mockDelete(ctx, id)
	}
	return nil
}
