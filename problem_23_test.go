package projecteuler

import "testing"

func TestIsPerfectNumber(t *testing.T) {
	AssertInt(IsPerfectNumber(28), 0, t)
	AssertInt(IsPerfectNumber(12), 1, t)
	for n := 1; n < 12; n++ {
		AssertTrue(IsPerfectNumber(n) != 1, t)
	}
}

func TestFindNotSumOfAbundants(t *testing.T) {
	AssertInt(
		FindNotSumOfAbundants(23),
		1+2+3+4+5+6+7+8+9+10+11+12+13+14+15+16+17+18+19+20+21+22+23,
		t)
	AssertInt(
		FindNotSumOfAbundants(24),
		1+2+3+4+5+6+7+8+9+10+11+12+13+14+15+16+17+18+19+20+21+22+23,
		t)
}

func TestProblem23(t *testing.T) {
	t.Parallel()
	AssertInt(Problem23(), 4179871, t)
}
