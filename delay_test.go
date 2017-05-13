package easy_middleware

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestDelay(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(Delay()(testHandler))
	defer test.Close()

	req, err := http.NewRequest(http.MethodGet, test.URL, nil)
	if err != nil {
		t.Errorf("Delay middleware: While creating the request is error occured (%s)", err)
		return
	}
	req.Header.Add("X-Add-Delay", "300ms")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Delay middleware request: %s", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Error("Delay middleware request: Unexpected good request")
	}
}
