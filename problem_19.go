package projecteuler

/*
  http://projecteuler.net/problem=19

  You are given the following information, but you may prefer to do some research for yourself.

  * 1 Jan 1900 was a Monday.
  * Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
  * A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

  How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

type Date struct {
	day   int
	month int // January == 0
	year  int
	dow   int // Sunday == 0
}

const (
	JANUARY = iota
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var daysInMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}

func NextDate(date Date) Date {
	date.dow = (date.dow + 1) % 7
	date.day++
	if date.day > daysInMonths[date.month] {
		if !(date.month == FEBRUARY && date.day == 29 && IsLeapYear(date.year)) {
			date.day = 1
			date.month = (date.month + 1) % 12
			if date.day == 1 && date.month == JANUARY {
				date.year++
			}
		}
	}

	return date
}

func Problem19() int {
	// We know 1 Jan 1900 is Monday; find date (with dow) for 1 Jan 1901.
	date := Date{1, JANUARY, 1900, MONDAY}
	for date.day != 1 || date.month != JANUARY || date.year != 1901 {
		date = NextDate(date)
	}

	// Now count all the first-of-month Sundays.  Include 31 Dec 2000.
	count := 0
	for date.day != 1 || date.month != JANUARY || date.year != 2001 {
		if date.day == 1 && date.dow == SUNDAY {
			count++
		}
		date = NextDate(date)
	}

	return count
}
