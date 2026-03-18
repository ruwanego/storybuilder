package router

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
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

	// create huma api
	api := humachi.New(r, huma.DefaultConfig("StoryBuilder API", "0.0.1"))

	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes using huma
	apiController.RegisterRoutes(api)
	sampleController.RegisterRoutes(api)

	return r
}
