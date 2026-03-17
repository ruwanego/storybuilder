package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/entities"
	"github.com/storybuilder/storybuilder/domain/usecases/sample"
	"github.com/storybuilder/storybuilder/transport/http/request"
	"github.com/storybuilder/storybuilder/transport/http/request/unpackers"
	"github.com/storybuilder/storybuilder/transport/http/response"
	"github.com/storybuilder/storybuilder/transport/http/response/transformers"
	"github.com/storybuilder/storybuilder/transport/http/errors"
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

// Get handles retrieving a list of samples.
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) error {
	// get the context
	ctx := r.Context()
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Get")
	// get data
	samples, err := ctl.sampleUseCase.Get(ctx)
	if err != nil {
		return err
	}
	// transform
	tr, err := response.Transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		return err
	}
	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, tr)
	return nil
}

// GetByID handles retrieving a single sample.
func (ctl *SampleController) GetByID(w http.ResponseWriter, r *http.Request) error {
	// get the context
	ctx := r.Context()
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.GetByID")
	// get id from request
	idVal := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idVal)
	// validate
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		return errors.ValidationMapError(errs)
	}
	// get data
	smpl, err := ctl.sampleUseCase.GetByID(ctx, id)
	if err != nil {
		return err
	}
	// transform
	tr, err := response.Transform(smpl, transformers.NewSampleTransformer(), false)
	if err != nil {
		return err
	}
	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, tr)
	return nil
}

// Add adds a new sample entry.
func (ctl *SampleController) Add(w http.ResponseWriter, r *http.Request) error {
	// get the context
	ctx := r.Context()
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Add")
	// unpack request
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		return err
	}
	// validate unpacked data
	errs := ctl.validator.Validate(sampleUnpacker)
	if errs != nil {
		return errors.ValidationMapError(errs)
	}
	// bind unpacked data to entities
	smpl := entities.Sample{
		Name:     sampleUnpacker.Name,
		Password: sampleUnpacker.Password,
	}
	// add
	err = ctl.sampleUseCase.Add(ctx, smpl)
	if err != nil {
		return err
	}
	// transform
	// tr := response.Transform(sample, transformers.NewSampleTransformer(), false)
	// send response
	ctl.sendResponse(ctx, w, http.StatusCreated)
	return nil
}

// Edit updates an existing sample entry.
func (ctl *SampleController) Edit(w http.ResponseWriter, r *http.Request) error {
	// get the context
	ctx := r.Context()
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Edit")
	// unpack request
	sampleUnpacker := unpackers.NewSampleUnpacker()
	err := request.Unpack(r, sampleUnpacker)
	if err != nil {
		return err
	}
	// get id from request
	idVal := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idVal)
	// validate request parameters
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		return errors.ValidationMapError(errs)
	}
	// validate unpacked data
	errs = ctl.validator.Validate(sampleUnpacker)
	if errs != nil {
		return errors.ValidationMapError(errs)
	}
	// bind unpacked data to entities
	smpl := entities.Sample{
		ID:       id,
		Name:     sampleUnpacker.Name,
		Password: sampleUnpacker.Password,
	}
	// edit
	err = ctl.sampleUseCase.Edit(ctx, smpl)
	if err != nil {
		return err
	}
	// send response
	ctl.sendResponse(ctx, w, http.StatusNoContent)
	return nil
}

// Delete deletes an existing sample entry.
func (ctl *SampleController) Delete(w http.ResponseWriter, r *http.Request) error {
	// get the context
	ctx := r.Context()
	// add a trace string to the context
	ctx = ctl.withTrace(ctx, "SampleController.Delete")
	// get id from request
	idVal := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idVal)
	// validate request parameters
	errs := ctl.validator.ValidateField(id, "required,gt=0")
	if errs != nil {
		return errors.ValidationMapError(errs)
	}
	// delete
	err := ctl.sampleUseCase.Delete(ctx, id)
	if err != nil {
		return err
	}
	// send response
	ctl.sendResponse(ctx, w, http.StatusNoContent)
	return nil
}
