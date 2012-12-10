package projecteuler

import "testing"

func TestFindSmallestNumber(t *testing.T) {
	AssertInt(FindSmallestNumber(10), 2520, t)
}

func TestProblem5(t *testing.T) {
	AssertInt(Problem5(), 232792560, t)
}
