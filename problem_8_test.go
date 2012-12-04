package projecteuler

import "testing"

func TestProblem8(t *testing.T) {
	t.Parallel()
	AssertInt(Problem8(), 40824, t)
}
