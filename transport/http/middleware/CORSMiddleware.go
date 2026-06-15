package middleware

import (
	"net/http"
	"strings"

	"github.com/storybuilder/storybuilder/app/config"
)

// CORSMiddleware attaches metrics to the request.
type CORSMiddleware struct {
	allowedOrigins []string
}

// NewCORSMiddleware returns a new instance of CORSMiddleware.
func NewCORSMiddleware(cfg config.AppConfig) *CORSMiddleware {
	return &CORSMiddleware{
		allowedOrigins: cfg.AllowedOrigins,
	}
}

func (m CORSMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedOrigin := ""

		if len(m.allowedOrigins) == 1 && m.allowedOrigins[0] == "*" {
			allowedOrigin = "*"
		} else {
			for _, o := range m.allowedOrigins {
				if o == origin {
					allowedOrigin = origin
					break
				}
				if strings.HasSuffix(o, "*") && strings.HasPrefix(origin, o[:len(o)-1]) {
					allowedOrigin = origin
					break
				}
			}
		}

		if allowedOrigin != "" {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Link")
			w.Header().Set("Access-Control-Max-Age", "300")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
