package easy_middleware_test

import (
	"github.com/donutloop/easy-middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoCache(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	testHandler := http.HandlerFunc(handler)

	test := httptest.NewServer(easy_middleware.NoCache()(testHandler))
	defer test.Close()

	response, err := http.Get(test.URL + "?limit=10")
	if err != nil {
		t.Errorf("url query middleware request: %s", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Error("url query middleware request: Unexpected bad request")
	}
}
