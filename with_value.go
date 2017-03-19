package easy_middleware

import (
	"net/http"
	"context"
)

// WithValue is a middleware that sets a given key/value in the request context.
func WithValue(key interface{}, val interface{}) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), key, val))
			h.ServeHTTP(w, r)
		})
	}
}
