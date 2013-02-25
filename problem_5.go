package projecteuler

/*
  http://projecteuler.net/problem=5

  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func FindSmallestNumber(limit int64) int64 {
	result := int64(1)

	for prime := range Primes() {
		if prime > limit {
			break
		}
		multiple := int64(1)
		for multiple < limit {
			multiple *= prime
		}
		result *= multiple / prime
	}

	return result
}

func Problem5() int64 {
	return FindSmallestNumber(20)
}
