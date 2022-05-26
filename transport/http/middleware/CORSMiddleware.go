package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CORSMiddleware attaches metrics to the request.
type CORSMiddleware struct{}

// NewCORSMiddleware returns a new instance of CORSMiddleware.
func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (m CORSMiddleware) Middleware(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})(next)
}
