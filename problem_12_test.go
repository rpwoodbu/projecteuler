package projecteuler

import "testing"

func TestTriangleNumbers(t *testing.T) {
	testNumbers := []int64{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}

	c := TriangleNumbers()
	for _, x := range testNumbers {
		AssertInt64(<-c, x, t)
	}
}

type DivisorCase struct {
	candidate int64
	divisors  []int64
}

func TestNumDivisors(t *testing.T) {
	cases := []DivisorCase{
		{1, []int64{1}},
		{3, []int64{1, 3}},
		{6, []int64{1, 2, 3, 6}},
		{10, []int64{1, 2, 5, 10}},
		{15, []int64{1, 3, 5, 15}},
		{21, []int64{1, 3, 7, 21}},
		{28, []int64{1, 2, 4, 7, 14, 28}},
		{36, []int64{1, 2, 3, 4, 6, 9, 12, 18, 36}},
	}

	for _, c := range cases {
		AssertInt(NumDivisors(c.candidate), len(c.divisors), t)
	}
}

func TestFirstToHaveNumDivisors(t *testing.T) {
	t.Parallel()
	result, ok := FirstToHaveNumDivisors(5)
	if ok {
		AssertInt64(result, 28, t)
	} else {
		t.Error("Returned not ok.")
	}
}

func TestProblem12(t *testing.T) {
	t.Parallel()
	result, ok := Problem12()
	if ok {
		AssertInt64(result, 76576500, t)
	} else {
		t.Error("Returned not ok.")
	}
}
