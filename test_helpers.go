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

func AssertTrue(got bool, t *testing.T) {
	if got != true {
		t.Fatalf("Expected true, got false")
	}
}

func AssertFalse(got bool, t *testing.T) {
	if got != false {
		t.Fatalf("Expected false, got true")
	}
}

func AssertIntArray(got []int, expected []int, t *testing.T) {
	if len(got) != len(expected) {
		t.Fatalf("Expected len %v, got len %v", len(expected), len(got))
	} else {
		for i := range got {
			if got[i] != expected[i] {
				t.Fatalf("Expected idx %v to be %v, got %v", i, expected[i], got[i])
				break
			}
		}
	}
}
