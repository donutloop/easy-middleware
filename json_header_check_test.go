package easy_middlware

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"bytes"
)

func TestJsonHeaderCheckMissingHeader(t *testing.T) {

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(JsonHeaderCheck(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))

	client := &http.Client{}
	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
	}

	if response.StatusCode != http.StatusUnsupportedMediaType {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %s)", response.StatusCode)
	}
}

func TestJsonHeaderCheckCorrectHeader(t *testing.T) {

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(JsonHeaderCheck(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	defer response.Body.Close()

	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %v)", response.StatusCode)
	}
}
