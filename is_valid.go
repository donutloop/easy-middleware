package easy_middleware

import (
	"net/http"
)

type validator interface {
	ok(w http.ResponseWriter, r http.Request) (bool, error)
}

func isValid(v validator) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if ok, err := v.ok(w, *r); !ok {
				err := NewError(err.Error(), http.StatusBadRequest)
				WriteJsonError(w, err, http.StatusBadRequest)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
