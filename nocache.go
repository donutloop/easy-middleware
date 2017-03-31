package easy_middleware

import (
	"net/http"
	"time"
)

// Unix epoch time
var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent
// a router from being cached by an upstream proxy and/or client.
func NoCache() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Delete any ETag headers that may have been set
			for _, v := range etagHeaders {
				if r.Header.Get(v) != "" {
					r.Header.Del(v)
				}
			}

			// Set our NoCache headers
			for k, v := range noCacheHeaders {
				w.Header().Set(k, v)
			}
			// server request
			h.ServeHTTP(w, r)
		})
	}
}
