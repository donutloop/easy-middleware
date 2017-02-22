package easy_middlware

import (
	"net/http"

	"encoding/json"
	"fmt"
)

type HttpError struct {
	Error string
}

func WriteJsonError(w http.ResponseWriter, v HttpError, code int)  (int, error) {
	b, err := json.Marshal(v)
	w.WriteHeader(code)
	if err != nil {
		// Fallback: if something going wrong while json marshal then use plain text error output
		return w.Write([]byte(fmt.Sprintf("Error: %s", v.Error)))
	}
	w.Header().Set("Content-Type", "application/json")
	return w.Write(b)
}
