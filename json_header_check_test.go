package easy_middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHeaderCheckMissingHeader(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(JsonHeaderCheck()(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
		return
	}
	defer response.Body.Close()

	herr := &ErrorResponse{}
	if err := json.NewDecoder(response.Body).Decode(herr); err != nil {
		t.Errorf("Json header check marschal body content: %s", err.Error())
		return
	}

	if response.StatusCode != http.StatusUnsupportedMediaType || herr.Error.Message != "Bad Content-Type or charset, expected 'application/json'" {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %s)", response.StatusCode)
	}
}

func TestJsonHeaderCheckCorrectHeader(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(JsonHeaderCheck()(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %v)", response.StatusCode)
	}
}
