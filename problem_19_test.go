package projecteuler

import "testing"

func TestIsLeapYear(t *testing.T) {
	AssertTrue(IsLeapYear(1904), t)
	AssertFalse(IsLeapYear(1905), t)
	AssertTrue(IsLeapYear(2000), t)
	AssertFalse(IsLeapYear(1900), t)
}

func TestNextDate(t *testing.T) {
	AssertTrue(NextDate(
		Date{31, DECEMBER, 1999, SATURDAY}) == Date{1, JANUARY, 2000, SUNDAY}, t)
	AssertTrue(NextDate(
		Date{28, FEBRUARY, 1992, MONDAY}) == Date{29, FEBRUARY, 1992, TUESDAY}, t)
	AssertTrue(NextDate(
		Date{28, FEBRUARY, 1995, WEDNESDAY}) == Date{1, MARCH, 1995, THURSDAY}, t)
}

func TestProblem19(t *testing.T) {
	AssertInt(Problem19(), 171, t)
}
