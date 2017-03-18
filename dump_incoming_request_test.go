package easy_middleware

import (
	"bytes"
	"net/http/httptest"
	"net/http"
	"testing"
	"log"
	"strings"
)

func TestDumpIncomingRequest(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	var b bytes.Buffer
	logger := log.New(&b, "", 0)

	testHandler := http.HandlerFunc(handler)
	server := httptest.NewServer(DumpIncomingRequest(logger)(testHandler))
	defer server.Close()

	response, err := http.Get(server.URL)

	if err != nil {
		t.Errorf("DumpIncomingRequest middleware request: %s", err.Error())
	}
	defer response.Body.Close()

	if !strings.Contains(b.String(), "User-Agent") {
		t.Errorf("Format of request is diffrent (%s)", b.String())
	}
}
