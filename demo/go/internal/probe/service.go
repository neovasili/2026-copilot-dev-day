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

// HttpClient keeps networking replaceable in tests and demos.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Service probes HTTP endpoints with timeout and bounded concurrency.
type Service struct {
	client  HttpClient
	timeout time.Duration
}

type probeJob struct {
	index int
	url   string
}

// NewService creates a new Service with the given HTTP client and timeout.
func NewService(client HttpClient, timeout time.Duration) *Service {
	if timeout <= 0 {
		timeout = 2 * time.Second
	}

	return &Service{
		client:  client,
		timeout: timeout,
	}
}

// Timeout returns the configured timeout for probes.
func (service *Service) Timeout() time.Duration {
	return service.timeout
}

// ProbeOnce performs a single probe to the given URL and returns the result.
func (service *Service) ProbeOnce(ctx context.Context, rawURL string) (Result, error) {
	if err := validateURL(rawURL); err != nil {
		wrapped := fmt.Errorf("invalid url %q: %w", rawURL, err)
		return Result{URL: rawURL, Err: wrapped}, wrapped
	}

	requestCtx, cancel := context.WithTimeout(ctx, service.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(requestCtx, http.MethodGet, rawURL, nil)
	if err != nil {
		wrapped := fmt.Errorf("build request for %q: %w", rawURL, err)
		return Result{URL: rawURL, Err: wrapped}, wrapped
	}

	start := time.Now()
	resp, err := service.client.Do(req)
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

// ProbeAll performs probes to all given URLs with the specified concurrency level.
func (service *Service) ProbeAll(ctx context.Context, urls []string, workers int) ([]Result, error) {
	if len(urls) == 0 {
		return []Result{}, nil
	}

	workers = normalizeWorkerCount(workers, len(urls))

	jobs := make(chan probeJob)
	results := make([]Result, len(urls))

	var (
		waitGroup sync.WaitGroup
		mutex     sync.Mutex
		firstErr  error
	)

	waitGroup.Add(workers)
	for workerIndex := 0; workerIndex < workers; workerIndex++ {
		go service.probeWorker(ctx, jobs, results, &firstErr, &mutex, &waitGroup)
	}

	for index, rawURL := range urls {
		select {
		case <-ctx.Done():
			close(jobs)
			waitGroup.Wait()
			return results, combineProbeAllErrors(ctx.Err(), firstErr)
		case jobs <- probeJob{index: index, url: rawURL}:
		}
	}

	close(jobs)
	waitGroup.Wait()
	return results, firstErr
}

// normalizeWorkerCount ensures the worker count is between 1 and the number of URLs.
func normalizeWorkerCount(workers, urlCount int) int {
	if workers <= 0 {
		return 1
	}
	if workers > urlCount {
		return urlCount
	}

	return workers
}

// combineProbeAllErrors combines the context error with the first probe error, if any.
func combineProbeAllErrors(contextError, firstErr error) error {
	if firstErr != nil {
		return errors.Join(contextError, firstErr)
	}

	return contextError
}

// probeWorker is a worker function that processes probe jobs and stores results.
func (service *Service) probeWorker(
	ctx context.Context,
	jobs <-chan probeJob,
	results []Result,
	firstErr *error,
	mutex *sync.Mutex,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	for job := range jobs {
		result, err := service.ProbeOnce(ctx, job.url)
		if err != nil {
			result.URL = job.url
			result.Err = err
		}

		mutex.Lock()
		results[job.index] = result
		if err != nil && *firstErr == nil {
			*firstErr = err
		}
		mutex.Unlock()
	}
}

// validateURL checks if the provided URL is valid and uses a supported scheme.
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
