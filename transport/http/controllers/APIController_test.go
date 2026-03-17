package controllers

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
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
	action := ctl.GetInfo

	handler := ctl.Wrap(action)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var body map[string]any
	err := json.Unmarshal(rr.Body.Bytes(), &body)
	assert.NoError(t, err)

	// The response is wrapped under a "data" key by response.Map. The value is the single transformer object.
	data, ok := body["data"].(map[string]any)
	assert.True(t, ok, "body.data should be a JSON object")
	assert.Equal(t, "StoryBuilder", data["name"])
	assert.Equal(t, "v0.0.1", data["version"])
}
