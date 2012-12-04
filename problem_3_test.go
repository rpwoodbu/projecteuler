package projecteuler

import "testing"

func TestProblem3(t *testing.T) {
	t.Parallel()
	AssertInt64(Problem3(), 6857, t)
}
