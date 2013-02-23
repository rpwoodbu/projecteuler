package projecteuler

import "strconv"
import "testing"

func AddNumbersHelper(t *testing.T, numbers ...int) {
	sum := 0
	var lines []string
	for _, n := range numbers {
		lines = append(lines, strconv.Itoa(n))
		sum += n
	}

	x := AddBigStringNumbers(lines)
	y, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	if y != sum {
		t.Logf("Got %s, expected %d", x, sum)
		t.Logf("  Inputs: %v", numbers)
		t.Fail()
	}
}

func TestAddBigNumbers(t *testing.T) {
	// These first tests make sure there aren't things like leading zeros.
	AssertIntArray(AddBigNumbers([][]int{[]int{1}, []int{1}}), []int{2}, t)
	AssertIntArray(AddBigNumbers([][]int{[]int{9}, []int{1}}), []int{1, 0}, t)

	AddNumbersHelper(t, 4587632, 3248767)
	AddNumbersHelper(t, 4587632, 3248767, 1234567)
	AddNumbersHelper(t, 9999999, 1000001)
}

func TestNumberConversions(t *testing.T) {
	AssertString(IntArrayToString([]int{1, 2, 3, 4, 5}), "12345", t)
	AssertIntArray(StringToIntArray("12345"), []int{1, 2, 3, 4, 5}, t)
}

func TestProblem13(t *testing.T) {
	t.Parallel()
	expected := "5537376230"
	answer := Problem13()
	if expected != answer {
		t.Errorf("Got %v, expected %v.", answer, expected)
	}
}
