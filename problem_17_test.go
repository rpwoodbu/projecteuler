package projecteuler

import "testing"

func TestNumberLetterCount(t *testing.T) {
	AssertInt(NumberLetterCount(16), 7, t)
	AssertInt(NumberLetterCount(20), 6, t)
	AssertInt(NumberLetterCount(21), 6+3, t)
	AssertInt(NumberLetterCount(999), 4+7+3+6+4, t)
	AssertInt(NumberLetterCount(100), 3+7, t)
	AssertInt(NumberLetterCount(500), 4+7, t)
	AssertInt(NumberLetterCount(1000), 3+8, t)
	AssertInt(NumberLetterCount(342), 23, t)
	AssertInt(NumberLetterCount(115), 20, t)
}

func TestNumbersLetterCount(t *testing.T) {
	AssertInt(NumbersLetterCount(5), 19, t)
}

func TestProblem17(t *testing.T) {
	t.Parallel()
	AssertInt(Problem17(), 21124, t)
}
