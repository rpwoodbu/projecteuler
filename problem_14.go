package projecteuler

/*
  http://projecteuler.net/problem=14

  The following iterative sequence is defined for the set of positive integers:

  n -> n/2 (n is even)
  n -> 3n + 1 (n is odd)

  Using the rule above and starting with 13, we generate the following sequence:

  13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

  It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

  Which starting number, under one million, produces the longest chain?

  NOTE: Once the chain starts the terms are allowed to go above one million.
*/

func Collatz(start int) chan int64 {
	c := make(chan int64)

	go func() {
		n := int64(start)
		for n != 1 {
			c <- n
			if n%2 == 0 { // Even
				n /= 2
			} else { // Odd
				n = (3 * n) + 1
			}
		}

		c <- 1
		close(c)
	}()

	return c
}

func Problem14() int {
	biggestCount := 0
	var biggestCountStart int

	for start := 1; start < 1000000; start++ {
		count := 0
		for _ = range Collatz(start) {
			count++
		}
		if count > biggestCount {
			biggestCount = count
			biggestCountStart = start
		}
	}

	return biggestCountStart
}
