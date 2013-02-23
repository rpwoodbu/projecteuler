package projecteuler

import "testing"

func TestBigFibonacci(t *testing.T) {
	terms := []string{
		"1",
		"1",
		"2",
		"3",
		"5",
		"8",
		"13",
		"21",
		"34",
		"55",
		"89",
		"144",
	}

	fibs := BigFibonacci()
	for i, term := range terms {
		r := IntArrayToString(<-fibs)
		if r != term {
			t.Errorf("BigFibonacci(%v) returned %v, expected %v", i+1, r, term)
		}
	}
}

func TestProblem25(t *testing.T) {
	t.Parallel()
	AssertInt(Problem25(), 4782, t)
}
