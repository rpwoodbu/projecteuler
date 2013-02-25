package projecteuler

import "testing"

// TODO: Consider removing this test in favor of TestFiniteUnitFraction.
func TestUnitFraction(t *testing.T) {
	type Test struct {
		denominator int
		DecimalFraction
	}

	tests := []Test{
		Test{2, DecimalFraction{[]int{0, 5}, nil}},
		Test{3, DecimalFraction{[]int{0}, []int{3, 3, 3}}},
		Test{4, DecimalFraction{[]int{0, 2, 5}, nil}},
		Test{5, DecimalFraction{[]int{0, 2}, nil}},
		Test{6, DecimalFraction{[]int{0, 1}, []int{6, 6}}},
		Test{7, DecimalFraction{[]int{0}, []int{1, 4, 2, 8, 5, 7}}},
		Test{8, DecimalFraction{[]int{0, 1, 2, 5}, nil}},
		Test{9, DecimalFraction{[]int{0}, []int{1, 1}}},
		Test{10, DecimalFraction{[]int{0, 1}, nil}},
	}

	for _, test := range tests {
		c := UnitFraction(test.denominator)

		TestDigits := func(digits []int) QuotientRemainder {
			var input QuotientRemainder
			var ok bool
			for _, digit := range digits {
				input, ok = <-c
				if !ok {
					t.Errorf("%d: Channel closed unexpectedly.", test.denominator)
					break
				}
				if input.q != digit {
					t.Errorf("%d: %d does not match expected digit %d.",
						test.denominator, input.q, digit)
					break
				}
			}
			return input
		}

		TestDigits(test.leadingDigits)
		if test.repeatingDigits != nil {
			a := TestDigits(test.repeatingDigits)
			b := TestDigits(test.repeatingDigits)
			if a != b {
				t.Errorf(
					"%d: Final quotient/remainder pair of repeating sequences not identical (%dr%d vs %dr%d)",
					test.denominator, a.q, a.r, b.q, b.r)
			}
		} else {
			_, ok := <-c
			if ok {
				t.Errorf("%d: Kept getting digits past end of finite sequence.",
					test.denominator)
			}
		}
	}
}

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
