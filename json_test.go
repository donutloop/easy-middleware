package easy_middlware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJson(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"echo": "test"})
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(Json()(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL)
	if err != nil {
		t.Errorf("Json middleware request: %s", err.Error())
	}
	defer response.Body.Close()

	if header := response.Header.Get("Content-Type"); header != "application/json" {
		t.Errorf("Json middleware request: %s", err.Error())
	}
}
