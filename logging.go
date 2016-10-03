package easy_middlware

import (
	"net/http"
	"time"
	"log"
)

// Logging of device request time
func Logging(logger *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			defer trace(logger, r)()
			h.ServeHTTP(w, r)
		})
	}
}

func trace(logger *log.Logger, r *http.Request) func() {
	start := time.Now()
	logger.Printf("Method: %s, url: %s, agent: %s started", r.Method, r.URL.Path, r.UserAgent())
	return func() {
		logger.Printf("Method: %s, url: %s, agent: %s completed in %v", r.Method, r.URL.Path, r.UserAgent() ,time.Since(start))
	}
}