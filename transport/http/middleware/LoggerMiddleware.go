package middleware

import (
	"net/http"
	"time"
)

// LoggerMiddleware attaches basic logging to the request.
type LoggerMiddleware struct{}

// NewLoggerMiddleware returns a new instance of LoggerMiddleware.
func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// standard logger middleware, passing to next
		next.ServeHTTP(w, r)
		// We could log `r.Method` and `r.URL.Path` and `time.Since(start)`
		_ = start
	})
}
