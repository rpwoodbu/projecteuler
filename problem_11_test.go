package projecteuler

import "testing"

func TestProblem11(t *testing.T) {
	t.Parallel()
	AssertInt(Problem11(), 70600674, t)
}
