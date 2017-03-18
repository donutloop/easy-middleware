package easy_middleware

import (
	"net/http"
	"log"
	"net/http/httputil"
)

//Dump request middleware creates a dump of in coming request.
func DumpIncomingRequest(logger *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := httputil.DumpRequest(r, true)
			if err == nil {
				logger.Println(string(b))
			}
			h.ServeHTTP(w, r)
		})
	}
}