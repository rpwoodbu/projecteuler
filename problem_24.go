package projecteuler

/*
	http://projecteuler.net/problem=24

	A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

	012   021   102   120   201   210

	What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

func Permute(place int, used []bool, sofar []byte, c chan []byte, closer bool) {
	if place < 0 {
		permutation := make([]byte, len(sofar))
		copy(permutation, sofar)
		c <- permutation
		return
	}

	for i := 0; i < len(used); i++ {
		if !used[i] {
			used[i] = true
			Permute(place-1, used, append(sofar, byte(i)), c, false)
			used[i] = false
		}
	}

	if closer {
		close(c)
	}
}

func Permutations(numValues int) chan []byte {
	c := make(chan []byte)

	go Permute(numValues-1, make([]bool, numValues), make([]byte, 0, numValues), c, true)

	return c
}

func Problem24() (string, bool) {
	count := 1
	for p := range Permutations(10) {
		if count == 1000000 {
			result := ""
			for _, i := range p {
				result += string(i + '0')
			}
			return result, true
		}
		count++
	}

	return "", false
}
