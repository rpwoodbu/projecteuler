package projecteuler

import "testing"

func TestProblem6(t *testing.T) {
	t.Parallel()
	AssertInt(Problem6(), 25164150, t)
}
