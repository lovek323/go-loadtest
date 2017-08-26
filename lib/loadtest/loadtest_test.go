package loadtest

import (
	"fmt"
	"testing"
)

func TestErrorIsNil(t *testing.T) {
	err := LoadTest(func() error {
		return nil
	})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestErrorIsReturned(t *testing.T) {
	err := LoadTest(func() error {
		return fmt.Errorf("Some error")
	})
	if err == nil {
		t.Errorf("Expected Some error")
	}
}
