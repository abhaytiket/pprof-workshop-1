package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleRequest(b *testing.B) {
	// Initialize any required global state, e.g., userCache
	userCache = NewUserCache()

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		b.Fatal(err)
	}

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleRequest)

		// Each iteration, we call our handler function, which is the one we're benchmarking.
		handler.ServeHTTP(rr, req)
	}
}
