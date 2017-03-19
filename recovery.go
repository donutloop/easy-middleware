package easy_middleware

import (
	"net/http"
	"net/http/httputil"
)

// Recovery middleware for panic
func Recovery(callback func(RequestDump []byte))  Middleware  {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					if b, err := httputil.DumpRequest(r, true); err == nil {
						callback(b)
					}
					WriteJsonError(w, NewError(err, http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}()

			// server request
			h.ServeHTTP(w, r)
		})
	}
}
