package easy_middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBodyLimitExceeded(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(SetBodyLimit(1)(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
	}
	defer response.Body.Close()

	herr := &ErrorResponse{}
	if err := json.NewDecoder(response.Body).Decode(herr); err != nil {
		t.Errorf("Json header check marschal body content: %s", err.Error())
	}

	if response.StatusCode != http.StatusBadRequest || herr.Error.Message != "Request body limit exceeded" {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %v)", response.StatusCode)
	}
}

func TestBodyLimit(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(SetBodyLimit(5000)(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %v)", response.StatusCode)
	}
}
