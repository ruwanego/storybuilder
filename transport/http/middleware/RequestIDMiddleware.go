package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// RequestIDMiddleware attaches metrics to the request.
type RequestIDMiddleware struct{}

// NewRequestIDMiddleware returns a new instance of RequestIDMiddleware.
func NewRequestIDMiddleware() *RequestIDMiddleware {
	return &RequestIDMiddleware{}
}

func (m RequestIDMiddleware) Middleware(next http.Handler) http.Handler {
	return middleware.RequestID(next)
}
