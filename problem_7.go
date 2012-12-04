package projecteuler

/*
	http://projecteuler.net/problem=7

	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

	What is the 10 001st prime number?
*/

func Problem7() int64 {
	primes := []int64{2}

	for x := int64(3); len(primes) < 10001; x++ {
		i := 0

		for ; i < len(primes); i++ {
			if x%primes[i] == 0 {
				break
			}
		}

		if i == len(primes) {
			// Got to end of list... must be prime.
			primes = append(primes, x)
		}
	}

	return primes[len(primes)-1]
}
