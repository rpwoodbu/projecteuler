package projecteuler

/*
	http://projecteuler.net/problem=7

	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

	What is the 10 001st prime number?
*/

func Primes() chan int64 {
	c := make(chan int64)

	go func() {
		c <- 2
		primes := []int64{}

	candidate:
		for x := int64(3); ; x += 2 {
			for _, prime := range primes {
				if x%prime == 0 {
					continue candidate
				}
				if prime*prime > x {
					// Past sqrt of the value, it'll never work.
					break
				}
			}

			// Got to end of list... must be prime.
			c <- x
			primes = append(primes, x)
		}
	}()

	return c
}

func Problem7() int64 {
	primes := Primes()
	for x := 0; x < 10000; x++ {
		<-primes
	}
	return <-primes
}
