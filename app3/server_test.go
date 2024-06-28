package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	//runtime.GOMAXPROCS(4)

	r := server()
	ts := httptest.NewServer(r)
	defer ts.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Get(ts.URL)
			if err != nil {
				b.Fatalf("could not send GET request: %v", err)
			}
			resp.Body.Close()
		}
	})
}
