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

func CollatzCounter(start int64, cache []int) int {
	count := 1

	n := start
	for n != 1 {
		if n < int64(len(cache)) {
			if cached := cache[n]; cached > 0 {
				count += cached
				break
			}
		}
		count++
		if n%2 == 0 { // Even
			n /= 2
		} else { // Odd
			n = (3 * n) + 1
		}
	}

	if start < int64(len(cache)) {
		cache[start] = count
	}
	return count
}

// This function isn't used by the problem, but is included
// to demonstrate the correctness of the algorithm, and is
// tested by the unit tests.
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

func Problem14() int64 {
	biggestCount := 0
	var biggestCountStart int64
	cache := make([]int, 1000000)

	for start := int64(1); start < 1000000; start++ {
		count := CollatzCounter(start, cache)
		if count > biggestCount {
			biggestCount = count
			biggestCountStart = start
		}
	}

	return biggestCountStart
}
