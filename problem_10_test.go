package projecteuler

import "testing"

func TestProblem10(t *testing.T) {
	t.Parallel()
	AssertInt64(Problem10(), 142913828922, t)
}
