package projecteuler

import "testing"

func TestProblem1(t *testing.T) {
	t.Parallel()
	AssertInt(Problem1(), 233168, t)
}
