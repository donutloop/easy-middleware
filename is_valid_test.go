package easy_middlware

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"bytes"
	"errors"
	"encoding/json"
)

type testValidator struct{}

var validatorOkBodyFunc func() (bool, error)
func (v testValidator) ok(w http.ResponseWriter, r http.Request) (bool, error){
	return validatorOkBodyFunc()
}

func TestIsValidFail(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	validatorOkBodyFunc = func() (bool, error){
		return false, errors.New("Something went wrong")
	}
	test := httptest.NewServer(isValid(testValidator{})(testHandler))
	defer test.Close()

	req, err := http.NewRequest("POST", test.URL, bytes.NewBuffer([]byte(`{"echo":"test"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Json header check request: %s", err.Error())
	}
	defer response.Body.Close()

	herr := &HttpError{}
	if err := json.NewDecoder(response.Body).Decode(herr); err != nil{
		t.Errorf("Json header check marschal body content: %s", err.Error())
	}

	if response.StatusCode != http.StatusBadRequest || herr.Error != "Something went wrong" {
		t.Errorf("Json middleware request: Header check isn't correct (StatusCode: %v)", response.StatusCode)
	}
}

func TestIsValidSuccess(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	validatorOkBodyFunc = func() (bool, error){
		return true, nil
	}
	test := httptest.NewServer(isValid(testValidator{})(testHandler))
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

