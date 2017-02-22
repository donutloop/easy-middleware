package easy_middlware

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"
)

// Set the Content-Type header 'application/json'
func Json() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jrw := &JsonResponseWriter{w}
			h.ServeHTTP(jrw, r)
		})
	}
}

// It implements the following interfaces:
// ResponseWriter
// http.ResponseWriter
// http.Flusher
// http.CloseNotifier
// http.Hijacker
type JsonResponseWriter struct {
	http.ResponseWriter
}

// Replace the parent EncodeJson to provide indentation.
func (w *JsonResponseWriter) EncodeJson(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (w *JsonResponseWriter) WriteError(herr interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := w.EncodeJson(herr)

	if err != nil {
		return err
	}

	w.WriteHeader(code)
	_, err = w.Write(b)
	return err
}

func (w *JsonResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Make sure the local EncodeJson and local Write are called.
// Does not call the parent WriteJson.
func (w *JsonResponseWriter) WriteJson(v interface{}) error {
	w.ResponseWriter.Header().Set("Content-Type", "application/json")
	b, err := w.EncodeJson(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// Call the parent WriteHeader.
func (w *JsonResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

// Make sure the local WriteHeader is called, and call the parent Flush.
// Provided in order to implement the http.Flusher interface.
func (w *JsonResponseWriter) Flush() {
	flusher := w.ResponseWriter.(http.Flusher)
	flusher.Flush()
}

// Call the parent CloseNotify.
// Provided in order to implement the http.CloseNotifier interface.
func (w *JsonResponseWriter) CloseNotify() <-chan bool {
	notifier := w.ResponseWriter.(http.CloseNotifier)
	return notifier.CloseNotify()
}

// Provided in order to implement the http.Hijacker interface.
func (w *JsonResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

// Make sure the local WriteHeader is called, and call the parent Write.
// Provided in order to implement the http.ResponseWriter interface.
func (w *JsonResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
