package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/copilot-dev-caceres/go-probe-demo/internal/probe"
)

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			w.WriteHeader(http.StatusOK)
		case "/slow":
			time.Sleep(75 * time.Millisecond)
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	svc := probe.NewService(testServer.Client(), 100*time.Millisecond)
	urls := []string{
		testServer.URL + "/ok",
		testServer.URL + "/slow",
		testServer.URL + "/missing",
		"not-a-url",
	}

	results, err := svc.ProbeAll(context.Background(), urls, 2)
	fmt.Printf("batch error: %v\n", err)

	for _, result := range results {
		fmt.Printf("url=%-30s status=%3d duration=%s err=%v\n", result.URL, result.StatusCode, result.Duration.Truncate(time.Millisecond), result.Err)
	}
}
