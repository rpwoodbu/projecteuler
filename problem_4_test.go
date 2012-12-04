package projecteuler

import "testing"

func TestProblem4(t *testing.T) {
	t.Parallel()
	AssertInt(Problem4(), 906609, t)
}
