package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	r := server()
	ts := httptest.NewServer(r)
	defer ts.Close()

	for i := 0; i < b.N; i++ {
		resp, err := http.Get(ts.URL)
		if err != nil {
			b.Fatalf("could not send GET request: %v", err)
		}
		resp.Body.Close()
	}
}
