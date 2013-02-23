package projecteuler

/*
	http://projecteuler.net/problem=25

	The Fibonacci sequence is defined by the recurrence relation:

	Fn = Fn1 + Fn2, where F1 = 1 and F2 = 1.
	Hence the first 12 terms will be:

	F1 = 1
	F2 = 1
	F3 = 2
	F4 = 3
	F5 = 5
	F6 = 8
	F7 = 13
	F8 = 21
	F9 = 34
	F10 = 55
	F11 = 89
	F12 = 144

	The 12th term, F12, is the first term to contain three digits.

	What is the first term in the Fibonacci sequence to contain 1000 digits?
*/

func BigFibonacci() <-chan []int {
	c := make(chan []int)

	go func() {
		c <- []int{1}
		c <- []int{1}

		a, b := []int{1}, []int{1}

		for i := 2; ; i++ {
			if len(a) < len(b) {
				// It will always be "a" that is short, and only by one digit.
				a = append([]int{0}, a...)
			}
			sum := AddBigNumbers([][]int{a, b})
			c <- sum
			a = b
			b = sum
		}
	}()

	return c
}

func Problem25() int {
	fibs := BigFibonacci()
	for i := 1; ; i++ {
		fib := <-fibs
		if len(fib) == 1000 {
			return i
		}
	}

	panic("Reached unreachable code.")
}
