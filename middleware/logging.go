// Package middleware provides HTTP middleware for request handling.
package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// Logging logs each incoming HTTP request with method, path, remote address, and response time.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote", r.RemoteAddr).
			Dur("duration", duration).
			Msg("Request handled")
	})
}
