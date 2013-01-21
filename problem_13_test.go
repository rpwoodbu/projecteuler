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

	x := AddBigNumbers(lines)
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
	AddNumbersHelper(t, 4587632, 3248767)
	AddNumbersHelper(t, 4587632, 3248767, 1234567)
	AddNumbersHelper(t, 9999999, 1000001)
}

func TestProblem13(t *testing.T) {
	t.Parallel()
	expected := "5537376230"
	answer := Problem13()
	if expected != answer {
		t.Fatalf("Got %v, expected %v.", answer, expected)
	}
}
