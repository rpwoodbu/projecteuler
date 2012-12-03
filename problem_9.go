package main

/*
  http://projecteuler.net/problem=9

  A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

  a^2 + b^2 = c^2
  For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

  There exists exactly one Pythagorean triplet for which a + b + c = 1000.
  Find the product abc.
*/

import "fmt"

func main() {
	for a := 0; a < 334; a++ {
		for b := a + 1; b < 1000-b-a-1; b++ {
			c := 1000 - b - a
			if a*a+b*b == c*c {
				fmt.Printf("%v^2 + %v^2 = %v^2\n", a, b, c)
				fmt.Printf("abc = %v\n", a*b*c)
				return
			}
		}
	}
}
