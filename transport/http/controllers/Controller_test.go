package controllers

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/globals"
)

func TestController_Wrap_Success(t *testing.T) {
	// Create mock container
	var logBuf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuf, nil))
	
	ctr := &container.Container{
		Adapters: container.Adapters{
			LogAdapter: logger,
		},
	}
	
	ctl := NewController(ctr)
	
	// Create an action that succeeds
	action := func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return nil
	}
	
	handler := ctl.Wrap(action)
	
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	
	handler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "OK", rr.Body.String())
}

func TestController_Wrap_Error(t *testing.T) {
	var logBuf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logBuf, nil))
	
	ctr := &container.Container{
		Adapters: container.Adapters{
			LogAdapter: logger,
		},
	}
	
	ctl := NewController(ctr)
	
	// Create an action that fails
	action := func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("something went wrong")
	}
	
	handler := ctl.Wrap(action)
	
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	
	handler.ServeHTTP(rr, req)
	
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "something went wrong")
}

func TestController_withTrace(t *testing.T) {
	ctl := &Controller{}
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := ctl.withTrace(req.Context(), "TestPrefix")
	val := ctx.Value(globals.PrefixKey)

	assert.NotNil(t, val)
	assert.Equal(t, "TestPrefix", val.(string))
}
