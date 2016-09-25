package easy_middlware

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestJson(t *testing.T) {

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"echo":"test"})
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(Json()(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("Json middleware request: %s", err.Error())
	}

	if header := response.Header.Get("Content-Type"); header != "application/json"{
		t.Errorf("Json middleware request: %s", err.Error())
	}
}
