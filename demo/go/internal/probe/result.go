package probe

import "time"

// Result captures one probe attempt.
type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}
