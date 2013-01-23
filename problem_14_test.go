package projecteuler

import "testing"

func CollatzHelper(expected []int64, t *testing.T) {
	c := Collatz(int(expected[0]))
	for _, n := range expected {
		answer := <-c
		if n != answer {
			t.Logf("Got %v, expected %v.", answer, n)
			t.Fail()
		}
	}
}

func TestCollatz(t *testing.T) {
	CollatzHelper([]int64{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}, t)
	CollatzHelper([]int64{1}, t) // Edge case.
}

func TestCollatzCounter(t *testing.T) {
	AssertInt(CollatzCounter(13), 10, t)
	AssertInt(CollatzCounter(1), 1, t) // Edge case.
}

func TestProblem14(t *testing.T) {
	t.Parallel()
	expected := int64(837799)
	answer := Problem14()
	if expected != answer {
		t.Errorf("Got %v, expected %v.", answer, expected)
	}
}
