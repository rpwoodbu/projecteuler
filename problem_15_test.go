package projecteuler

import "testing"

func TestGridRouteCounter(t *testing.T) {
	AssertInt64(GridRouteCounter(2, 2), 6, t)
	AssertInt64(GridRouteCounter(3, 3), 20, t)
}

func TestProblem15(t *testing.T) {
	t.Parallel()
	AssertInt64(Problem15(), 137846528820, t)
}
