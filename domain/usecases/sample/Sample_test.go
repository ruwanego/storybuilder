package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
)

func TestNewSample(t *testing.T) {
	ctr := &container.Container{
		Repositories: container.Repositories{
			SampleRepository: &mockSampleRepository{},
		},
	}
	sample := NewSample(ctr)
	assert.NotNil(t, sample)
}
