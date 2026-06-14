package router

import (
	"net/http"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/transport/http/controllers"
	"github.com/storybuilder/storybuilder/transport/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container, appCfg config.AppConfig) http.Handler { // will fix config later
	mux := http.NewServeMux()

	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	mux.Handle("GET /{$}", apiController.Wrap(apiController.GetInfo))

	mux.HandleFunc("GET /openapi", serveSwaggerUI)
	mux.HandleFunc("GET /openapi/swagger.yaml", serveSwaggerSpec)

	mux.Handle("GET /samples", sampleController.Wrap(sampleController.Get))
	mux.Handle("GET /samples/{id}", sampleController.Wrap(sampleController.GetByID))
	mux.Handle("POST /samples", sampleController.Wrap(sampleController.Add))
	mux.Handle("PUT /samples/{id}", sampleController.Wrap(sampleController.Edit))
	mux.Handle("DELETE /samples/{id}", sampleController.Wrap(sampleController.Delete))

	// add middleware to router
	var handler http.Handler = mux
	middlewares := middleware.Init(ctr, appCfg)
	// Apply middlewares in reverse order to ensure correct execution order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
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
