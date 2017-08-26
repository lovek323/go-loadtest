package loadtest

import (
	"time"
)

type Summary struct {
	Taken time.Duration
}

func LoadTest(function func() error) (*Summary, error) {
	start := time.Now()
	if err := function(); err != nil {
		return nil, err
	}
	taken := time.Since(start)
	summary := &Summary{
		Taken: taken,
	}
	return summary, nil
}
