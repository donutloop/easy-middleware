package easy_middlware

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestThenOrdersHandlersCorrectly(t *testing.T) {

	middlewareBase := func(tag string) Middleware {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tag))
				h.ServeHTTP(w, r)
			})
		}
	}

	t1 := middlewareBase("t1\n")
	t2 := middlewareBase("t2\n")
	t3 := middlewareBase("t3\n")

	testEndpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("endpoint\n"))
	})

	chained := New(t1, t2, t3).Then(testEndpoint)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	chained.ServeHTTP(w, r)

	if w.Body.String() != "t1\nt2\nt3\nendpoint\n" {
		t.Errorf("Then does not order handlers correctly (Order: %s)", w.Body.String())
	}
}

func TestCreateOrdersHandlersCorrectly(t *testing.T) {

	middlewareBase := func(tag string) Middleware {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tag))
				h.ServeHTTP(w, r)
			})
		}
	}

	t1 := middlewareBase("t1\n")
	t2 := middlewareBase("t2\n")
	t3 := middlewareBase("t3\n")

	testEndpoint := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("endpoint\n"))
	})

	chained := New(t1, t2, t3)

	t4 := middlewareBase("t4\n")
	createdChained := Create(chained, t4).Then(testEndpoint)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	createdChained.ServeHTTP(w, r)

	if w.Body.String() != "t1\nt2\nt3\nt4\nendpoint\n" {
		t.Errorf("Then does not order handlers correctly (Order: %s)", w.Body.String())
	}
}