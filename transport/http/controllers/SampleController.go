package controllers

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/entities"
	"github.com/storybuilder/storybuilder/domain/usecases/sample"
	"github.com/storybuilder/storybuilder/transport/http/response"
	"github.com/storybuilder/storybuilder/transport/http/response/transformers"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	*Controller
	sampleUseCase *sample.Sample
}

// NewSampleController creates a new instance of the controller.
func NewSampleController(ctr *container.Container) *SampleController {
	return &SampleController{
		Controller:    NewController(ctr),
		sampleUseCase: sample.NewSample(ctr),
	}
}

// RegisterRoutes binds the controller routes to the Huma API.
func (ctl *SampleController) RegisterRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-samples",
		Method:      http.MethodGet,
		Path:        "/samples",
		Summary:     "Get samples list",
		Description: "Returns a collection of samples",
		Tags:        []string{"samples"},
	}, ctl.Get)

	huma.Register(api, huma.Operation{
		OperationID: "get-sample-by-id",
		Method:      http.MethodGet,
		Path:        "/samples/{id}",
		Summary:     "Get sample by ID",
		Description: "Returns a single sample resource",
		Tags:        []string{"samples"},
	}, ctl.GetByID)

	huma.Register(api, huma.Operation{
		OperationID: "add-sample",
		Method:      http.MethodPost,
		Path:        "/samples",
		Summary:     "Create a new sample",
		Description: "Creates a sample resource",
		Tags:        []string{"samples"},
		DefaultStatus: http.StatusCreated,
	}, ctl.Add)

	huma.Register(api, huma.Operation{
		OperationID: "edit-sample",
		Method:      http.MethodPut,
		Path:        "/samples/{id}",
		Summary:     "Update a sample",
		Description: "Updates a sample resource",
		Tags:        []string{"samples"},
		DefaultStatus: http.StatusNoContent,
	}, ctl.Edit)

	huma.Register(api, huma.Operation{
		OperationID: "delete-sample",
		Method:      http.MethodDelete,
		Path:        "/samples/{id}",
		Summary:     "Delete a sample",
		Description: "Deletes a sample resource",
		Tags:        []string{"samples"},
		DefaultStatus: http.StatusNoContent,
	}, ctl.Delete)
}

// Types for Handlers

type GetSamplesOutput struct {
	Body struct {
		Data []transformers.SampleTransformer `json:"data"`
	}
}

type GetSampleByIDInput struct {
	ID int `path:"id" required:"true" minimum:"1" doc:"Sample ID"`
}

type GetSampleByIDOutput struct {
	Body struct {
		Data transformers.SampleTransformer `json:"data"`
	}
}

// Huma uses native tags like required:"true" replacing our previous validation logic
type SamplePayload struct {
	Name     string `json:"name" required:"true"`
	Password string `json:"password" required:"true"`
}

type AddSampleInput struct {
	Body SamplePayload
}

type EditSampleInput struct {
	ID   int `path:"id" required:"true" minimum:"1" doc:"Sample ID"`
	Body SamplePayload
}

type DeleteSampleInput struct {
	ID int `path:"id" required:"true" minimum:"1" doc:"Sample ID"`
}

// Get handles retrieving a list of samples.
func (ctl *SampleController) Get(ctx context.Context, input *struct{}) (*GetSamplesOutput, error) {
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Get")
	// get data
	samples, err := ctl.sampleUseCase.Get(ctx)
	if err != nil {
		return nil, err
	}
	// transform
	tr, err := response.Transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		return nil, err
	}

	resp := &GetSamplesOutput{}
	resp.Body.Data = tr.([]transformers.SampleTransformer)

	return resp, nil
}

// GetByID handles retrieving a single sample.
func (ctl *SampleController) GetByID(ctx context.Context, input *GetSampleByIDInput) (*GetSampleByIDOutput, error) {
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.GetByID")
	// get data
	smpl, err := ctl.sampleUseCase.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	// transform
	tr, err := response.Transform(smpl, transformers.NewSampleTransformer(), false)
	if err != nil {
		return nil, err
	}

	resp := &GetSampleByIDOutput{}
	resp.Body.Data = tr.(transformers.SampleTransformer)

	return resp, nil
}

// Add adds a new sample entry.
func (ctl *SampleController) Add(ctx context.Context, input *AddSampleInput) (*struct{}, error) {
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Add")

	// bind unpacked data to entities
	smpl := entities.Sample{
		Name:     input.Body.Name,
		Password: input.Body.Password,
	}
	// add
	err := ctl.sampleUseCase.Add(ctx, smpl)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(ctx context.Context, input *EditSampleInput) (*struct{}, error) {
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Edit")

	// bind unpacked data to entities
	smpl := entities.Sample{
		ID:       input.ID,
		Name:     input.Body.Name,
		Password: input.Body.Password,
	}
	// edit
	err := ctl.sampleUseCase.Edit(ctx, smpl)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete deletes an existing sample entry.
func (ctl *SampleController) Delete(ctx context.Context, input *DeleteSampleInput) (*struct{}, error) {
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Delete")

	// delete
	err := ctl.sampleUseCase.Delete(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
