package projecteuler

/*
  http://projecteuler.net/problem=22

	Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

	For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 x 53 = 49714.

	What is the total of all the name scores in the file?
*/

import (
	"io"
	"os"
	"sort"
)

func LoadNames(filename string) (names []string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1024)
	inQuotes := false
	currentName := ""
	for {
		count, err := f.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for _, c := range data[:count] {
			if c == '"' {
				inQuotes = !inQuotes
				if !inQuotes {
					names = append(names, currentName)
					currentName = ""
				}
			} else if inQuotes {
				currentName += string(c)
			}
		}
	}

	return
}

func AlphaValue(name string) (value int) {
	for _, c := range name {
		value += int(c - 'A' + 1)
	}
	return
}

func LoadNamesSorted(filename string) (names []string) {
	names = LoadNames(filename)
	sort.Strings(names)
	return
}

func Problem22() (sum int) {
	names := LoadNamesSorted("data/names.txt")
	for i, name := range names {
		sum += (i + 1) * AlphaValue(name)
	}
	return
}
