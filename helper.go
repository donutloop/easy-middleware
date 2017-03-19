package easy_middleware

import (
	"net/http"

	"encoding/json"
	"fmt"
)

func NewError(message interface{}, code int ) ErrorResponse {
	return ErrorResponse{
		ErrorEntity{
			Message: message,
			Code: code,
		},
	}
}

type ErrorResponse struct {
	Error ErrorEntity `json:"error"`
}

type ErrorEntity struct {
	// Code is the HTTP response status code and will always be populated.
	Code int `json:"code"`
	Message interface{} `json:"message"`
}

func WriteJsonError(w http.ResponseWriter, v ErrorResponse, code int) (int, error) {
	b, err := json.Marshal(v)
	w.WriteHeader(code)
	if err != nil {
		// Fallback: if something going wrong while json marshal then use plain text error output
		return w.Write([]byte(fmt.Sprintf("Error: %s", v.Error.Message)))
	}
	w.Header().Set("Content-Type", "application/json")
	return w.Write(b)
}
