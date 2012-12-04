package projecteuler

/*
  http://projecteuler.net/problem=4

  A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.

  Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"sort"
	"strconv"
)

func IsPalindrome(s string) bool {
	length := len(s)
	for x := 0; x < length/2; x++ {
		if s[x] != s[length-x-1] {
			return false
		}
	}
	return true
}

func Problem4() int {
	palindromes := []int{}

	for x := 999; x > 0; x-- {
		for y := 999; y > 0; y-- {
			product := x * y
			if IsPalindrome(strconv.Itoa(product)) {
				palindromes = append(palindromes, product)
			}
		}
	}

	sort.Ints(palindromes)
	return palindromes[len(palindromes)-1]
}
