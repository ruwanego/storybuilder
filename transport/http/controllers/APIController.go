package controllers

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

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

// RegisterRoutes binds the controller routes to the Huma API.
func (ctl *APIController) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-info",
		Method:      http.MethodGet,
		Path:        "/",
		Summary:     "Get API information",
		Description: "Returns metadata about the running API (name, version, purpose)",
		Tags:        []string{"information"},
	}, ctl.GetInfo)
}

type GetInfoOutput struct {
	Body struct {
		Data transformers.APITransformer `json:"data"`
	}
}

// GetInfo returns basic details of the API.
func (ctl *APIController) GetInfo(ctx context.Context, input *struct{}) (*GetInfoOutput, error) {
	// transform
	tr := transformers.APITransformer{
		Name:    "StoryBuilder",
		Version: "v0.0.1",
		Purpose: "REST API base written in Golang",
	}

	resp := &GetInfoOutput{}
	resp.Body.Data = tr

	return resp, nil
}
