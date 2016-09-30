package easy_middlware

import (
	"bytes"
	"testing"
	"net/http"
	"encoding/json"
	"net/http/httptest"
	"log"
	"strings"
)

func TestLogging(t *testing.T) {

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"echo":"test"})
	}

	var b bytes.Buffer
	logger := log.New(&b, "", 0)

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(Logging(logger)(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL)
	defer response.Body.Close()

	if err != nil {
		t.Errorf("logging middleware request: %s", err.Error())
	}

	if  strings.Contains(b.String(), "Started GET Go-http-client/1.1 Completed in") {
		t.Errorf("logging middleware request: log output should match %q is a string", b.String())
	}
}



