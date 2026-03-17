package router

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
)

// newTestContainer builds a minimal container for router tests.
func newTestContainer() *container.Container {
	var logBuf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuf, nil))
	return &container.Container{
		Adapters: container.Adapters{
			LogAdapter: logger,
		},
	}
}

func TestInit_ReturnsRouter(t *testing.T) {
	ctr := newTestContainer()
	mux := Init(ctr)
	assert.NotNil(t, mux)
}

func TestRouter_GetInfo(t *testing.T) {
	ctr := newTestContainer()
	mux := Init(ctr)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRouter_MethodNotAllowed(t *testing.T) {
	ctr := newTestContainer()
	mux := Init(ctr)

	// POST to "/" which only has a GET handler — chi should return 405
	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestRouter_NotFound(t *testing.T) {
	ctr := newTestContainer()
	mux := Init(ctr)

	req, _ := http.NewRequest("GET", "/nonexistent-route", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
