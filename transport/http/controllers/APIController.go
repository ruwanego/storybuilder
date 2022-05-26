package controllers

import (
	"net/http"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/transport/http/response/transformers"
)

// APIController contains controller logic for endpoints.
type APIController struct {
	*Controller
}

// NewAPIController creates a new instance of the controller.
func NewAPIController(ctr *container.Container) *APIController {
	return &APIController{
		Controller: NewController(ctr),
	}
}

// GetInfo return basic details of the API.
func (ctl *APIController) GetInfo(w http.ResponseWriter, r *http.Request) {
	// transform
	tr := transformers.APITransformer{
		Name:    "StoryBuilder",
		Version: "v0.0.1",
		Purpose: "REST API base written in Golang",
	}
	// send response
	ctl.sendResponse(r.Context(), w, http.StatusOK, tr)
}
