package loadtest

import (
	"fmt"
	"testing"
	"time"
)

var doSomething = func() error {
	return nil
}
var doError = func() error {
	return fmt.Errorf("Some error")
}
var current = time.Unix(0, 0)
var incrementingTimeProvider = func() time.Time {
	current = current.Add(1 * time.Second)
	return current
}

func TestErrorIsReturned(t *testing.T) {
	now = incrementingTimeProvider
	_, err := LoadTest(doError, 1)
	if err == nil {
		t.Errorf("Expected Some error")
	}
}

func TestSummaryIsReturned(t *testing.T) {
	now = incrementingTimeProvider
	expectedStart := current.Add(1 * time.Second)
	summary, err := LoadTest(doSomething, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedEnd := current
	if summary.Start != expectedStart {
		t.Errorf("Expected %s, got %s", expectedStart, summary.Start)
	}
	if summary.End != expectedEnd {
		t.Errorf("Expected %s, got %s", expectedEnd, summary.End)
	}
}
