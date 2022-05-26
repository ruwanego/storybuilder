package middleware

import (
	"net/http"

	"github.com/storybuilder/storybuilder/app/container"
)

type Middleware interface {
	Middleware(next http.Handler) http.Handler
}

func Init(ctr *container.Container) (middlewares []func(http.Handler) http.Handler) {
	// NOTE: middleware will execute in the order they are added to the router
	middlewares = []func(http.Handler) http.Handler{
		// add metrics middleware first
		NewMetricsMiddleware().Middleware,
		NewCORSMiddleware().Middleware,
		NewLoggerMiddleware().Middleware,
		NewRequestIDMiddleware().Middleware,
		NewRequestCheckerMiddleware(ctr).Middleware,
		NewRequestAlterMiddleware().Middleware,
	}
	return
}
