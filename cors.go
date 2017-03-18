package easy_middleware

import (
	"net/http"
)

//CORS Middleware creates a new CORS Middleware with default headers.
func Cors(defaultHeaders  string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// If "api-cors-header" is not given, but "api-enable-cors" is true, we set cors to "*"
			// otherwise, all head values will be passed to HTTP handler
			corsHeaders := defaultHeaders
			if corsHeaders == "" {
				corsHeaders = "*"
			}

			w.Header().Add("Access-Control-Allow-Origin", corsHeaders)
			w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
			w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
			h.ServeHTTP(w, r)
		})
	}
}
