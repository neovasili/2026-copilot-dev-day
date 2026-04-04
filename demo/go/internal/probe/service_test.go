package probe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"
)

func TestProbeOnceInvalidURL(t *testing.T) {
	svc := NewService(http.DefaultClient, 250*time.Millisecond)

	result, err := svc.ProbeOnce(context.Background(), "not-a-url")
	if err == nil {
		t.Fatal("expected error for invalid URL")
	}
	if result.URL != "not-a-url" {
		t.Fatalf("expected URL to be preserved, got %q", result.URL)
	}
}

func TestProbeOnceReturnsStatusCode(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))
	defer testServer.Close()

	svc := NewService(testServer.Client(), 250*time.Millisecond)

	result, err := svc.ProbeOnce(context.Background(), testServer.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.StatusCode != http.StatusAccepted {
		t.Fatalf("expected status %d, got %d", http.StatusAccepted, result.StatusCode)
	}
	if result.Duration <= 0 {
		t.Fatalf("expected duration > 0, got %s", result.Duration)
	}
}

func TestProbeAllPreservesInputOrder(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/slow":
			time.Sleep(40 * time.Millisecond)
			w.WriteHeader(http.StatusNoContent)
		case "/fast":
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer testServer.Close()

	svc := NewService(testServer.Client(), 300*time.Millisecond)
	urls := []string{
		testServer.URL + "/slow",
		testServer.URL + "/fast",
		testServer.URL + "/missing",
	}

	results, err := svc.ProbeAll(context.Background(), urls, 3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(results) != len(urls) {
		t.Fatalf("expected %d results, got %d", len(urls), len(results))
	}

	for i := range urls {
		if results[i].URL != urls[i] {
			t.Fatalf("order mismatch at index %d: want %q, got %q", i, urls[i], results[i].URL)
		}
	}

	statusCodes := []int{results[0].StatusCode, results[1].StatusCode, results[2].StatusCode}
	sort.Ints(statusCodes)
	want := []int{http.StatusOK, http.StatusNoContent, http.StatusNotFound}
	for i := range want {
		if statusCodes[i] != want[i] {
			t.Fatalf("unexpected status set: got %v want %v", statusCodes, want)
		}
	}
}

func TestProbeAllReturnsErrorWhenAnyProbeFails(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	svc := NewService(testServer.Client(), 200*time.Millisecond)
	urls := []string{testServer.URL, "invalid-url"}

	results, err := svc.ProbeAll(context.Background(), urls, 2)
	if err == nil {
		t.Fatal("expected batch error")
	}
	if len(results) != len(urls) {
		t.Fatalf("expected %d results, got %d", len(urls), len(results))
	}
	if results[1].Err == nil {
		t.Fatal("expected per-item error on invalid URL")
	}
}

func TestProbeAllReturnsContextErrorWhenCancelled(t *testing.T) {
	blockingServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(300 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer blockingServer.Close()

	svc := NewService(blockingServer.Client(), 2*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	_, err := svc.ProbeAll(ctx, []string{blockingServer.URL, blockingServer.URL}, 1)
	if err == nil {
		t.Fatal("expected cancellation error")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context deadline error, got %v", err)
	}
}

func ExampleService_ProbeAll() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	svc := NewService(testServer.Client(), 200*time.Millisecond)
	results, err := svc.ProbeAll(context.Background(), []string{testServer.URL}, 1)
	fmt.Println(err == nil)
	fmt.Println(results[0].StatusCode)

	// Output:
	// true
	// 200
}
