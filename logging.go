package easy_middleware

import (
	"net/http"
	"time"
	"fmt"
)

// Logging of device request time
func Logging(callback func(s string)) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer trace(callback, r)()
			h.ServeHTTP(w, r)
		})
	}
}

func trace(callback func(s string), r *http.Request) func() {
	start := time.Now()
	callback(fmt.Sprintf("Method: %s, url: %s, agent: %s started", r.Method, r.URL.Path, r.UserAgent()))
	return func() {
		callback(fmt.Sprintf("Method: %s, url: %s, agent: %s completed in %v", r.Method, r.URL.Path, r.UserAgent(), time.Since(start)))
	}
}
