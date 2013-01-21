package projecteuler

import "testing"

func TestPowerTwoDigitSum(t *testing.T) {
	AssertInt(PowerTwoDigitSum(15), 26, t)
}

func TestProblem16(t *testing.T) {
	t.Parallel()
	AssertInt(Problem16(), 1366, t)
}
