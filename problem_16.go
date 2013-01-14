package projecteuler

/*
  http://projecteuler.net/problem=16

  2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

  What is the sum of the digits of the number 2^1000?
*/

func DoubleBase10(x []int) []int {
	result := []int{}
	carry := 0

	for _, i := range x {
		i = (i * 2) + carry
		if i >= 10 {
			carry = i / 10
			i = i % 10
		} else {
			carry = 0
		}
		result = append(result, i)
	}

	if carry != 0 {
		result = append(result, carry)
	}

	return result
}

func PowerTwoDigitSum(power int) int {
	value := []int{2}
	for x := 0; x < power-1; x++ {
		value = DoubleBase10(value)
	}

	sum := 0
	for _, x := range value {
		sum += x
	}

	return sum
}

func Problem16() int {
	return PowerTwoDigitSum(1000)
}
