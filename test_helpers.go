package projecteuler

import "testing"

func AssertInt(got int, expected int, t *testing.T) {
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func AssertInt64(got int64, expected int64, t *testing.T) {
	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}
