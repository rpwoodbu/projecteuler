package projecteuler

import "testing"

func TestSumOfDivisors(t *testing.T) {
	AssertInt(SumOfDivisors(1), 1, t)
	AssertInt(SumOfDivisors(4), 3, t)
	AssertInt(SumOfDivisors(220), 284, t)
	AssertInt(SumOfDivisors(284), 220, t)
}

func TestProblem21(t *testing.T) {
	t.Parallel()
	AssertInt(Problem21(), 31626, t)
}
