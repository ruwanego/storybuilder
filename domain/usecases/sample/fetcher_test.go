package sample

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/entities"
)

func TestSample_GetByID(t *testing.T) {
	mockRepo := &mockSampleRepository{
		mockGetByID: func(ctx context.Context, id int) (entities.Sample, error) {
			if id == 1 {
				return entities.Sample{ID: 1, Name: "Test"}, nil
			}
			return entities.Sample{}, nil
		},
	}

	ctr := &container.Container{
		Repositories: container.Repositories{
			SampleRepository: mockRepo,
		},
	}
	sample := NewSample(ctr)

	// Test found
	res, err := sample.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, res.ID)

	// Test not found (id == 0 returns errorNoSample)
	res2, err2 := sample.GetByID(context.Background(), 2)
	assert.Error(t, err2)
	assert.Equal(t, 0, res2.ID)
}
