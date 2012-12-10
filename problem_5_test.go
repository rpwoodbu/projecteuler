package projecteuler

import "testing"

func TestFindSmallestNumber(t *testing.T) {
	result, ok := FindSmallestNumber(10)
	if ok {
		AssertInt(result, 2520, t)
	} else {
		t.Fatal("Returned not ok.")
	}
}

func TestProblem5(t *testing.T) {
	t.Parallel()
	result, ok := Problem5()
	if ok {
		AssertInt(result, 232792560, t)
	} else {
		t.Fatal("Returned not ok.")
	}
}
