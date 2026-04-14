package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/copilot-dev-caceres/go-probe-demo/internal/probe"
	"github.com/copilot-dev-caceres/go-probe-demo/internal/tools"
)

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/ok":
			responseWriter.WriteHeader(http.StatusOK)
		case "/slow":
			time.Sleep(75 * time.Millisecond)
			responseWriter.WriteHeader(http.StatusNoContent)
		case "/timeout":
			time.Sleep(150 * time.Millisecond)
			responseWriter.WriteHeader(http.StatusNoContent)
		default:
			responseWriter.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	svc := probe.NewService(testServer.Client(), 100*time.Millisecond)
	urls := []string{
		testServer.URL + "/ok",
		testServer.URL + "/slow",
		testServer.URL + "/timeout",
		testServer.URL + "/missing",
		"not-a-url",
		"ftp://wrong-scheme-url",
	}
	fmt.Println(
		tools.Magenta(fmt.Sprintf("\nProbing %d urls with concurrency 2 and timeout %s\n", len(urls), svc.Timeout())),
	)

	results, err := svc.ProbeAll(context.Background(), urls, 2)
	if err != nil {
		fmt.Println(tools.Red(fmt.Sprintf("Batch error: %v", err)))
	}

	for _, result := range results {
		if result.Err != nil {
			fmt.Println(
				tools.Red(
					fmt.Sprintf("url=%-30s status=%3d duration=%s err=%v", result.URL, result.StatusCode, result.Duration.Truncate(time.Millisecond), result.Err),
				),
			)
			continue
		}
		fmt.Println(
			tools.Green(
				fmt.Sprintf("url=%-30s status=%3d duration=%s", result.URL, result.StatusCode, result.Duration.Truncate(time.Millisecond)),
			),
		)
	}
}
