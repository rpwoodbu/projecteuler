package projecteuler

/*
  http://projecteuler.net/problem=5

  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func FindSmallestNumber(limit int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	result := 1

	for _, prime := range primes {
		if prime > limit {
			break
		}
		multiple := 1
		for multiple < limit {
			multiple *= prime
		}
		result *= multiple / prime
	}

	return result
}

func Problem5() int {
	return FindSmallestNumber(20)
}
