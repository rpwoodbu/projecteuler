package main

/*
  http://projecteuler.net/problem=10

  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

  Find the sum of all the primes below two million.
*/

import "fmt"

const TARGET = 2000000

func main() {
	primes := []int64{2}

	for x := int64(3); x < TARGET; x += 2 {
		i := 1 // Skip first prime (2) as we are never checking evens.

		for ; i < len(primes); i++ {
			if x%primes[i] == 0 {
				break
			}
			if primes[i]*2 > x {
				// Past half the value, it'll never work.
				i = len(primes) // Bail, signaling a prime.
				break
			}
		}

		if i == len(primes) {
			// Got to end of list... must be prime.
			primes = append(primes, x)
			if len(primes)%100 == 0 {
				fmt.Printf("%v (%v)\r", x, len(primes))
			}
		}
	}
	fmt.Println()

	sum := int64(0)
	for i := 0; i < len(primes); i++ {
		sum += primes[i]
	}

	fmt.Println(sum)
}
