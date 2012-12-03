package main

/*
  http://projecteuler.net/problem=5

  2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import "fmt"

func main() {
outer_loop:
	for x := 1; ; x++ {
		for y := 2; y <= 20; y++ {
			if x%y != 0 {
				continue outer_loop
			}
		}
		fmt.Println(x)
		break
	}
}
