package main

/*
  http://projecteuler.net/problem=2

  Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

  1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

  By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

import "fmt"

func Fibonacci() chan (int) {
	c := make(chan (int))

	go func() {
		x := 0
		y := 1

		for {
			sum := x + y
			c <- sum
			x = y
			y = sum
		}
	}()

	return c
}

func main() {
	sum := 0
	for x := range Fibonacci() {
		if x >= 4000000 {
			break
		}
		if x%2 == 0 {
			sum += x
		}
	}

	fmt.Println(sum)
}
