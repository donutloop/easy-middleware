package easy_middleware


import (
	"net/http"
	"time"
)

const DELAY_HEADER_KEY = "X-Add-Delay"

// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
func Delay() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			delayHeaderValue := r.Header.Get(DELAY_HEADER_KEY)
			if delayHeaderValue != "" {
				delayDuration, err := time.ParseDuration(delayHeaderValue)
				if err == nil {
					time.Sleep(delayDuration)
				}
			}

			h.ServeHTTP(rw, r)
		})
	}
}


