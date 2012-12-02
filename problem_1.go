package main

/*
 * http://projecteuler.net/problem=1
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

import "fmt"

func main() {
  sum := 0
  for x := 0; x < 1000; x++ {
    if x % 3 == 0 || x % 5 == 0 {
      sum += x
    }
  }

  fmt.Println(sum)
}
