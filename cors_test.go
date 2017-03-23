package easy_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	server := httptest.NewServer(Cors("")(testHandler))
	defer server.Close()

	req, err := http.NewRequest("Get", server.URL, nil)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Cors check request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Access-Control-Allow-Origin is bad (Expected: %s, Acutal: %s)", "*", response.Header.Get("Access-Control-Allow-Origin"))
	}

	if response.Header.Get("Access-Control-Allow-Headers") != "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth" {
		t.Errorf("Access-Control-Allow-Headers is bad (Expected: %s, Acutal: %s)", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth", response.Header.Get("Access-Control-Allow-Headers"))
	}

	if response.Header.Get("Access-Control-Allow-Methods") != "HEAD, GET, POST, DELETE, PUT, OPTIONS" {
		t.Errorf("Access-Control-Allow-Methods is bad (Expected: %s, Acutal: %s)", "HEAD, GET, POST, DELETE, PUT, OPTIONS", response.Header.Get("Access-Control-Allow-Methods"))
	}
}
