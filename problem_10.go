package projecteuler

/*
  http://projecteuler.net/problem=10

  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

  Find the sum of all the primes below two million.
*/

const PROB10_TARGET = 2000000

func Problem10() (sum int64) {
	for prime := range Primes() {
		if prime >= PROB10_TARGET {
			break
		}
		sum += prime
	}
	return
}
