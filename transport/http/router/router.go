package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/transport/http/controllers"
	"github.com/storybuilder/storybuilder/transport/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container) *chi.Mux {
	// create new router
	r := chi.NewRouter()
	// add middleware to router
	r.Use(middleware.Init(ctr)...)
	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)
	// bind controller functions to routes
	// api info
	r.Get("/", apiController.GetInfo)
	// sample
	r.Get("/samples", sampleController.Get)
	r.Get("/samples/{id:[0-9]+}", sampleController.GetByID)
	r.Post("/samples", sampleController.Add)
	r.Put("/samples/{id:[0-9]+}", sampleController.Edit)
	r.Delete("/samples/{id:[0-9]+}", sampleController.Delete)
	return r
}
