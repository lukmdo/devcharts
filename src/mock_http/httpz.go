package main

import (
	"net/http"
	"time"
	"log"
)

// ResponseWriter extends ResponseWriter to access Code
// lightweight httptest.ResponseRecorder
type ResponseWriter struct {
	http.ResponseWriter
	Code int
}

// WriteHeader wraps original call
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)
	rw.Code = code
}

// NewResponseWriter is constructor for http.ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, 200}
}

// log/debug middleware
func loggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := NewResponseWriter(w)

		now := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL, nw.Code, time.Since(now))
	})
}

