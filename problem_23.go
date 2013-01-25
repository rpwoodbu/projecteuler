package projecteuler

/*
	http://projecteuler.net/problem=23

	A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

	A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

	As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

	Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

// Returns 0 if number is perfect, -1 if deficient, and 1 if abundant.
func IsPerfectNumber(n int) int {
	sum := SumOfDivisors(n)
	switch {
	case sum == n:
		return 0
	case sum < n:
		return -1
	case sum > n:
		return 1
	}
	panic("Unreachable code reached.")
}

func FindNotSumOfAbundants(upperBound int) (sum int) {
	// Seed abundants[0] = false so that array index is the actual number.
	abundants := []bool{false}
	for n := 1; n <= upperBound; n++ {
		if IsPerfectNumber(n) == 1 {
			abundants = append(abundants, true)
		} else {
			abundants = append(abundants, false)
		}
	}

candidate:
	for n := 1; n <= upperBound; n++ {
		for i := 1; i < len(abundants); i++ {
			if !abundants[i] {
				continue
			}
			if i >= n {
				break
			}
			if abundants[n-i] {
				continue candidate
			}
		}
		// Got here, so cannot be expressed as sum of two abundant numbers.
		sum += n
	}

	return
}

func Problem23() int {
	return FindNotSumOfAbundants(28123)
}
