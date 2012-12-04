package projecteuler

import "testing"

func TestProblem2(t *testing.T) {
	t.Parallel()
	AssertInt(Problem2(), 4613732, t)
}
