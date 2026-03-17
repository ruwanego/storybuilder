package controllers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/globals"
	"github.com/storybuilder/storybuilder/transport/http/response"
)

// Controller is the base struct that holds fields and functionality common to all controllers.
type Controller struct {
	logger    *slog.Logger
	validator adapters.ValidatorAdapterInterface
}

// NewController creates a new instance of the controller.
func NewController(ctr *container.Container) *Controller {
	return &Controller{
		logger:    ctr.Adapters.LogAdapter,
		validator: ctr.Adapters.ValidatorAdapter,
	}
}

// withTrace adds an optional tracing string that will be displayed in error messages.
func (ctl *Controller) withTrace(ctx context.Context, prefix string) context.Context {
	return globals.AddTrace(ctx, prefix)
}

// sendResponse is a convenience function wrapping the actual `response.Send` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendResponse(_ context.Context, w http.ResponseWriter, code int, payload ...any) {
	if len(payload) == 0 {
		response.Send(w, nil, code)
		return
	}
	response.Send(w, response.Map(payload), code)
}

// sendError is a convenience function wrapping the actual `response.Error` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendError(ctx context.Context, w http.ResponseWriter, err any) {
	response.Error(ctx, w, err, ctl.logger)
}

// Action represents an HTTP handler that can return an error.
type Action func(w http.ResponseWriter, r *http.Request) error

// Wrap takes an Action and converts it to a standard http.HandlerFunc.
// If the Action returns an error, Wrap will automatically send the error response.
func (ctl *Controller) Wrap(action Action) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := action(w, r); err != nil {
			// use the request context to ensure tracing works
			ctl.sendError(r.Context(), w, err)
		}
	}
}
