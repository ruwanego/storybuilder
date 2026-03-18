package controllers

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
)

func TestAPIController_GetInfo(t *testing.T) {
	var logBuf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuf, nil))

	ctr := &container.Container{
		Adapters: container.Adapters{
			LogAdapter: logger,
		},
	}

	ctl := NewAPIController(ctr)

	// Test the Huma handler directly
	resp, err := ctl.GetInfo(context.Background(), &struct{}{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	data := resp.Body.Data
	assert.Equal(t, "StoryBuilder", data.Name)
	assert.Equal(t, "v0.0.1", data.Version)
}
