package projecteuler

/*
  http://projecteuler.net/problem=21

	Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
	If d(a) = b and d(b) = a, where a != b, then a and b are an amicable pair and each of a and b are called amicable numbers.

	For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

	Evaluate the sum of all the amicable numbers under 10000.
*/

func SumOfDivisors(n int) (sum int) {
	for x := n / 2; x > 0; x-- {
		if n%x == 0 {
			sum += x
		}
	}
	return
}

func Problem21() (sum int) {
	// Precalculating divisors provides about a 33% speedup.
	var divisors [10000]int

	for x := 1; x < len(divisors); x++ {
		divisors[x] = SumOfDivisors(x)
	}

	for x := 1; x < len(divisors); x++ {
		d1 := divisors[x]
		if x != d1 && d1 < len(divisors) {
			d2 := divisors[d1]
			if d2 == x {
				sum += x
			}
		}
	}

	return
}

// This is not used, but shows a simple solution.
func Problem21Simple() (sum int) {
	for x := 1; x < 10000; x++ {
		d := SumOfDivisors(x)
		if x != d && d < 10000 && x == SumOfDivisors(d) {
			sum += x
		}
	}
	return
}
