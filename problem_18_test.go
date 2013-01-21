package projecteuler

import "testing"

func TestMaxPath(t *testing.T) {
	triangle := [][]int{
		[]int{3},
		[]int{7, 4},
		[]int{2, 4, 6},
		[]int{8, 5, 9, 3},
	}

	AssertInt(MaxPath(triangle), 23, t)
}

func TestProblem18(t *testing.T) {
	t.Parallel()
	AssertInt(Problem18(), 1074, t)
}
