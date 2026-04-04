package probe

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// HTTPClient keeps networking replaceable in tests and demos.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Service probes HTTP endpoints with timeout and bounded concurrency.
type Service struct {
	client  HTTPClient
	timeout time.Duration
}

func NewService(client HTTPClient, timeout time.Duration) *Service {
	if timeout <= 0 {
		timeout = 2 * time.Second
	}

	return &Service{
		client:  client,
		timeout: timeout,
	}
}

func (s *Service) ProbeOnce(ctx context.Context, rawURL string) (Result, error) {
	if err := validateURL(rawURL); err != nil {
		wrapped := fmt.Errorf("invalid url %q: %w", rawURL, err)
		return Result{URL: rawURL, Err: wrapped}, wrapped
	}

	requestCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet, rawURL, nil)
	if err != nil {
		wrapped := fmt.Errorf("build request for %q: %w", rawURL, err)
		return Result{URL: rawURL, Err: wrapped}, wrapped
	}

	start := time.Now()
	resp, err := s.client.Do(req)
	duration := time.Since(start)
	if err != nil {
		wrapped := fmt.Errorf("probe %q: %w", rawURL, err)
		return Result{URL: rawURL, Duration: duration, Err: wrapped}, wrapped
	}
	defer resp.Body.Close()

	return Result{
		URL:        rawURL,
		StatusCode: resp.StatusCode,
		Duration:   duration,
	}, nil
}

func (s *Service) ProbeAll(ctx context.Context, urls []string, workers int) ([]Result, error) {
	if len(urls) == 0 {
		return []Result{}, nil
	}

	if workers <= 0 {
		workers = 1
	}
	if workers > len(urls) {
		workers = len(urls)
	}

	type job struct {
		index int
		url   string
	}

	jobs := make(chan job)
	results := make([]Result, len(urls))

	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		firstErr error
	)

	worker := func() {
		defer wg.Done()
		for j := range jobs {
			result, err := s.ProbeOnce(ctx, j.url)
			if err != nil {
				result.URL = j.url
				result.Err = err
			}

			mu.Lock()
			results[j.index] = result
			if err != nil && firstErr == nil {
				firstErr = err
			}
			mu.Unlock()
		}
	}

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker()
	}

	for i, rawURL := range urls {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			if firstErr != nil {
				return results, errors.Join(ctx.Err(), firstErr)
			}
			return results, ctx.Err()
		case jobs <- job{index: i, url: rawURL}:
		}
	}

	close(jobs)
	wg.Wait()
	return results, firstErr
}

func validateURL(rawURL string) error {
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return fmt.Errorf("scheme %q is not supported", parsed.Scheme)
	}

	if parsed.Host == "" {
		return errors.New("missing host")
	}

	return nil
}
