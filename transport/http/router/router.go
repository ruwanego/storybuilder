package router

import (
	"net/http"

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
	r.Get("/", apiController.Wrap(apiController.GetInfo))
	// docs (Swagger UI)
	r.Get("/openapi", serveSwaggerUI)
	r.Get("/openapi/swagger.yaml", serveSwaggerSpec)
	// sample
	r.Get("/samples", sampleController.Wrap(sampleController.Get))
	r.Get("/samples/{id:[0-9]+}", sampleController.Wrap(sampleController.GetByID))
	r.Post("/samples", sampleController.Wrap(sampleController.Add))
	r.Put("/samples/{id:[0-9]+}", sampleController.Wrap(sampleController.Edit))
	r.Delete("/samples/{id:[0-9]+}", sampleController.Wrap(sampleController.Delete))
	return r
}

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>StoryBuilder API Docs</title>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@4/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@4/swagger-ui-bundle.js" crossorigin></script>
  <script>
    window.onload = function () {
      SwaggerUIBundle({
        url: '/openapi/swagger.yaml',
        dom_id: '#swagger-ui',
        presets: [SwaggerUIBundle.presets.apis],
        layout: 'BaseLayout',
      });
    };
  </script>
</body>
</html>
`))
}

func serveSwaggerSpec(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/swagger.yaml")
}
