package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// LoggerMiddleware attaches metrics to the request.
type LoggerMiddleware struct{}

// NewLoggerMiddleware returns a new instance of LoggerMiddleware.
func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return middleware.Logger(next)
}
