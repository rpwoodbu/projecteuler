package projecteuler

import "testing"

func TestFiniteUnitFraction(t *testing.T) {
	type Test struct {
		denominator int
		DecimalFraction
	}

	tests := []Test{
		Test{2, DecimalFraction{[]int{0, 5}, nil}},
		Test{3, DecimalFraction{[]int{0}, []int{3}}},
		Test{4, DecimalFraction{[]int{0, 2, 5}, nil}},
		Test{5, DecimalFraction{[]int{0, 2}, nil}},
		Test{6, DecimalFraction{[]int{0, 1}, []int{6}}},
		Test{7, DecimalFraction{[]int{0}, []int{1, 4, 2, 8, 5, 7}}},
		Test{8, DecimalFraction{[]int{0, 1, 2, 5}, nil}},
		Test{9, DecimalFraction{[]int{0}, []int{1}}},
		Test{10, DecimalFraction{[]int{0, 1}, nil}},
	}

	for _, test := range tests {
		df := FiniteUnitFraction(test.denominator)
		AssertIntArray(df.leadingDigits, test.leadingDigits, t)
		AssertIntArray(df.repeatingDigits, test.repeatingDigits, t)
	}
}

func TestProblem26(t *testing.T) {
	t.Parallel()
	AssertInt(Problem26(), 983, t)
}
