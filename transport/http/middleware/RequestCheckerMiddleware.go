package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/transport/http/errors"
	"github.com/storybuilder/storybuilder/transport/http/response"
)

// RequestCheckerMiddleware checks if request configuration is in order.
type RequestCheckerMiddleware struct {
	omittedRoutes []string
	ctr           *container.Container
}

// NewRequestCheckerMiddleware returns a new instance of RequestCheckerMiddleware.
func NewRequestCheckerMiddleware(ctr *container.Container) *RequestCheckerMiddleware {
	return &RequestCheckerMiddleware{
		omittedRoutes: []string{"/openapi", "/metrics"},
		ctr:           ctr,
	}
}

// Middleware executes middleware rules of RequestCheckerMiddleware.
func (m *RequestCheckerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqURL := r.URL.String()
		for v := range slices.Values(m.omittedRoutes) {
			if strings.HasPrefix(reqURL, v) {
				next.ServeHTTP(w, r)
				return
			}
		}

		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				err := errors.NewMiddlewareError(
					fmt.Sprintf("API only accepts JSON as Content-Type, '%s' is given", contentType), 100, "")
				response.Error(r.Context(), w, err, m.ctr.Adapters.LogAdapter)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
