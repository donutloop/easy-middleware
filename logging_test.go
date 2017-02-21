package easy_middlware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogging(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"echo": "test"})
	}

	var b bytes.Buffer
	logger := log.New(&b, "", 0)

	testHandler := http.HandlerFunc(handler)
	test := httptest.NewServer(Logging(logger)(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL)

	if err != nil {
		t.Errorf("logging middleware request: %s", err.Error())
	}
	defer response.Body.Close()

	if strings.Contains(b.String(), "Completed in") && strings.Contains(b.String(), "started") {
		t.Errorf("logging middleware request: log output should match %q is a string", b.String())
	}
}
