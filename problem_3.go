package projecteuler

/*
  http://projecteuler.net/problem=3

  The prime factors of 13195 are 5, 7, 13 and 29.

  What is the largest prime factor of the number 600851475143 ?
*/

const PROB3_TARGET int64 = 600851475143

func Problem3() int64 {
	target := PROB3_TARGET
	var factor int64

	for x := int64(2); x < target; x++ {
		if target%x == 0 {
			factor = x
			target /= x
			x = 1
		}
	}

	if target != 1 {
		factor = target
	}

	return factor
}
