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

	req, err := http.NewRequest(http.MethodPost, test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	if err != nil {
		t.Errorf("Json header check request: %v", err)
		return
	}

	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err)
		return
	}
	defer response.Body.Close()

	herr := new(ErrorResponse)
	if err := json.NewDecoder(response.Body).Decode(herr); err != nil {
		t.Errorf("Json header check marschal body content: %s", err)
		return
	}

	if response.StatusCode != http.StatusUnsupportedMediaType || herr.Error.Message != "Bad Content-Type or charset, expected 'application/json'" {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %d)", response.StatusCode)
	}
}

func TestJsonHeaderCheckCorrectHeader(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(JsonHeaderCheck()(testHandler))
	defer test.Close()

	req, err := http.NewRequest(http.MethodPost, test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	if err != nil {
		t.Errorf("Json header check request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %d)", response.StatusCode)
	}
}
