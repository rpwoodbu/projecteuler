package projecteuler

import "testing"

func TestPrimes(t *testing.T) {
	primes := []int64{2, 3, 5, 7, 11, 13}
	c := Primes()
	for _, prime := range primes {
		AssertInt64(<-c, prime, t)
	}
}

func TestProblem7(t *testing.T) {
	t.Parallel()
	AssertInt64(Problem7(), 104743, t)
}
