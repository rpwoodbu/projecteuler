package projecteuler

import "testing"

func TestAddDigits(t *testing.T) {
	AssertInt(AddDigits([]int{1, 2, 3}), 6, t)
}

func TestMultiplyBigNumbers(t *testing.T) {
	AssertIntArray(MultiplyBigNumbers([][]int{{9}, {9}}), BigNumber(9*9), t)
	AssertIntArray(MultiplyBigNumbers([][]int{{9, 9}, {9}}), BigNumber(99*9), t)
	AssertIntArray(MultiplyBigNumbers([][]int{{9}, {9}, {9}}), BigNumber(9*9*9), t)
	AssertIntArray(MultiplyBigNumbers([][]int{{9, 9}, {9, 9}}), BigNumber(99*99), t)
	AssertIntArray(MultiplyBigNumbers([][]int{{3, 2, 1}, {6, 5, 4}}), BigNumber(123*456), t)
	AssertIntArray(MultiplyBigNumbers([][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}}),
		BigNumber(123*456*789), t)
}

func TestFactorial(t *testing.T) {
	f := Factorial(10)
	AssertIntArray(f, BigNumber(3628800), t)
}

func TestProblem20(t *testing.T) {
	AssertInt(Problem20(), 648, t)
}
