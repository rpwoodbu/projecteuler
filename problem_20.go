package projecteuler

/*
  http:projecteuler.net/problem=20

  n! means n x (n - 1) x ... x 3 x 2 x 1

  For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800,
  and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

  Find the sum of the digits in the number 100!
*/

func MultiplyBigNumbers(bigNums [][]int) []int {
	result := bigNums[0]

	for _, bigNum := range bigNums[1:] {
		accumulator := []int{0}
		for i := range bigNum {
			carry := 0
			for j := range result {
				for len(accumulator) < j+i+1 {
					accumulator = append(accumulator, 0)
				}
				a := accumulator[j+i] + (bigNum[i] * result[j]) + carry
				accumulator[j+i] = a % 10
				carry = a / 10
			}
			if carry != 0 {
				if len(accumulator) < len(result)+i+1 {
					accumulator = append(accumulator, carry)
				} else {
					accumulator[len(result)+i] = carry
				}
			}
		}
		result = accumulator
	}

	return result
}

func AddDigits(digits []int) int {
	sum := 0
	for _, digit := range digits {
		sum += digit
	}

	return sum
}

func BigNumber(n int) []int {
	result := []int{}
	for ; n != 0; n /= 10 {
		result = append(result, n%10)
	}
	return result
}

func Factorial(f int) []int {
	factors := [][]int{}
	for ; f > 0; f-- {
		factors = append(factors, BigNumber(f))
	}
	return MultiplyBigNumbers(factors)
}

func Problem20() int {
	return AddDigits(Factorial(100))
}
