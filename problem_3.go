package main

/*
  http://projecteuler.net/problem=3

  The prime factors of 13195 are 5, 7, 13 and 29.

  What is the largest prime factor of the number 600851475143 ?
*/

import "fmt"

const TARGET int64 = 600851475143

func main() {
	target := TARGET
	for x := int64(2); x < target; x++ {
		if target%x == 0 {
			fmt.Println(x)
			target /= x
			x = 1
		}
	}

	if target != 1 {
		fmt.Println(target)
	}
}
