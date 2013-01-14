package projecteuler

/*
  http://projecteuler.net/problem=17

  If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

  If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


  NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

var numbers = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"", // Don't use for 10-19.
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func NumberWordsBelow100(x int) string {
	if x >= 100 {
		panic("Only use for numbers below 100.")
	}

	if x < 20 {
		return numbers[x]
	}

	tensPlace := tens[((x%100)/10)-1]
	if x%10 == 0 {
		return tensPlace
	}

	return tensPlace + "-" + numbers[x%10]
}

func NumberWords(x int) string {
	if x > 1000 {
		panic("Cannot handle numbers greater than 1000.")
	}

	if x == 1000 {
		return "one thousand"
	}

	if x < 100 {
		return NumberWordsBelow100(x)
	}

	hundredsPlace := numbers[x/100]
	if x%100 == 0 {
		return hundredsPlace + " hundred"
	}

	return hundredsPlace + " hundred and " + NumberWordsBelow100(x%100)
}

func NumberLetterCount(x int) int {
	s := NumberWords(x)
	count := 0

	for _, c := range s {
		if c != ' ' && c != '-' {
			count++
		}
	}

	return count
}

func NumbersLetterCount(x int) int {
	sum := 0

	for i := 1; i <= x; i++ {
		sum += NumberLetterCount(i)
	}

	return sum
}

func Problem17() int {
	return NumbersLetterCount(1000)
}
