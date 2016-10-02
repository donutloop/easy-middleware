package easy_middlware

import (
	"net/http"
)

func SetBodyLimit(limit int64) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

			if r.ContentLength > limit {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

