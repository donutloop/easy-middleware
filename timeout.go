package easy_middleware

import (
	"context"
	"net/http"
	"time"
)

// Timeout is a middleware that cancels ctx after a given timeout and return
// a 504 Gateway Timeout error to the client.
func Timeout(timeoutAfter time.Duration) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			ctx, cancel := context.WithTimeout(r.Context(), timeoutAfter * time.Second)
			defer func() {
				cancel()
				if ctx.Err() == context.DeadlineExceeded {
					rw.WriteHeader(http.StatusGatewayTimeout)
				}
			}()

			r = r.WithContext(ctx)
			h.ServeHTTP(rw, r)
		})
	}
}
