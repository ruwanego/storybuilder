package sample

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
)

func TestSample_Delete(t *testing.T) {
	mockRepo := &mockSampleRepository{
		mockDelete: func(ctx context.Context, id int) error {
			if id == 1 {
				return nil
			}
			return errors.New("not found")
		},
	}

	ctr := &container.Container{
		Repositories: container.Repositories{
			SampleRepository: mockRepo,
		},
	}
	sample := NewSample(ctr)

	err := sample.Delete(context.Background(), 1)
	assert.NoError(t, err)

	err2 := sample.Delete(context.Background(), 2)
	assert.Error(t, err2)
}
