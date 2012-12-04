package projecteuler

import "testing"

func TestProblem7(t *testing.T) {
	t.Parallel()
	AssertInt64(Problem7(), 104743, t)
}
