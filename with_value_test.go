package easy_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithValueMiddleware(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		v := r.Context().Value("test")

		if v.(string) != "test" {
			t.Errorf("Unexpected value (%s)", v)
		}
	}

	server := httptest.NewServer(WithValue("test", "test")(http.HandlerFunc(testHandler)))
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("WithValue middleware request: %s", err.Error())
		return
	}
	defer response.Body.Close()
}
