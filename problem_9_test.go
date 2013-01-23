package projecteuler

import "testing"

func TestProblem9(t *testing.T) {
	t.Parallel()
	result, ok := Problem9()
	if ok {
		AssertInt(result, 31875000, t)
	} else {
		t.Error("Returned not ok.")
	}
}
