package easy_middlware

import (
	"net/http"
	"time"
	"log"
)

// Logging of device request time
func Logging(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		logger.Printf("Started %s %s", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)

		logger.Printf("%s Completed in %v", r.UserAgent() ,time.Since(start))
	})
}

