// Package middleware provides HTTP middleware for request handling.
package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

// Recovery intercepts panics during HTTP request handling and returns a 500 error,
// logging the stack trace and error details.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Error().
					Interface("error", rec).
					Bytes("stacktrace", debug.Stack()).
					Str("path", r.URL.Path).
					Msg("Panic recovered")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
