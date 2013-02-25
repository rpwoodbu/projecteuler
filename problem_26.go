package projecteuler

/*
	http://projecteuler.net/problem=26

	A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

	1/2  = 	0.5
	1/3  = 	0.(3)
	1/4  = 	0.25
	1/5  = 	0.2
	1/6  = 	0.1(6)
	1/7  = 	0.(142857)
	1/8  = 	0.125
	1/9  = 	0.(1)
	1/10 = 	0.1

	Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

	Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

type QuotientRemainder struct {
	q int
	r int
}

type DecimalFraction struct {
	leadingDigits   []int
	repeatingDigits []int
}

func UnitFraction(d int) <-chan QuotientRemainder {
	c := make(chan QuotientRemainder)

	go func() {
		qr := QuotientRemainder{0, 1}

		for {
			qr.q = qr.r / d
			qr.r = qr.r % d
			c <- qr
			if qr.r == 0 {
				break
			}
			qr.r *= 10
		}

		close(c)
	}()

	return c
}

// Figure out if there's a recurring cycle in a unit fraction, and express
// it in a finite way.
func FiniteUnitFraction(d int) (df DecimalFraction) {
	qrs := make(map[QuotientRemainder]int)
	digits := []int{}

	for qr := range UnitFraction(d) {
		if i, ok := qrs[qr]; ok {
			df.leadingDigits = digits[:i]
			df.repeatingDigits = digits[i:]
			return
		}
		qrs[qr] = len(digits)
		digits = append(digits, qr.q)
	}

	// If we got here, this is a finite decimal.
	df.leadingDigits = digits
	return
}

func Problem26() int {
	var longestCycle, longestCycleDenominator int
	for d := 3; d < 1000; d += 2 {
		df := FiniteUnitFraction(d)
		if len(df.repeatingDigits) > longestCycle {
			longestCycle = len(df.repeatingDigits)
			longestCycleDenominator = d
		}
	}
	return longestCycleDenominator
}
