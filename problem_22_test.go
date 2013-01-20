package projecteuler

import "testing"

func TestLoadNames(t *testing.T) {
	names := LoadNames("data/names.txt")
	AssertInt(len(names), 5163, t)
}

func TestLoadNamesSorted(t *testing.T) {
	names := LoadNamesSorted("data/names.txt")
	AssertInt(len(names), 5163, t)
	AssertTrue(names[937] == "COLIN", t)
}

func TestAlphaValue(t *testing.T) {
	AssertInt(AlphaValue("COLIN"), 53, t)
}

func TestProblem22(t *testing.T) {
	AssertInt(Problem22(), 871198282, t)
}
