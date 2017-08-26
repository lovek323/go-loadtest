package loadtest

import (
	"fmt"
	"testing"
)

var doSomething = func() error {
	return nil
}
var doError = func() error {
	return fmt.Errorf("Some error")
}

func TestErrorIsNil(t *testing.T) {
	_, err := LoadTest(doSomething)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestErrorIsReturned(t *testing.T) {
	_, err := LoadTest(doError)
	if err == nil {
		t.Errorf("Expected Some error")
	}
}
