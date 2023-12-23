package piscine

import "github.com/01-edu/z01"

func printNbrBase(nbr int, base string) {
	if notValid(base) {
		NV()
		return
	}

	num, isNegative := getCoordinates(nbr, base)

	printNum(num, base, isNegative)
}

func printNum(num []int, base string, isNegative bool) {
	if isNegative {
		printR('-')
	}

	for i := len(num) - 1; i >= 0; i-- {
		index := num[i]
		z01.PrintRune(rune(base[index]))
	}
}

func getCoordinates(nbr int, base string) ([]int, bool) {
	l := len(base)
	var r int
	var result []int
	var isNegative bool = false

	if nbr < 0 {
		isNegative = true
		nbr = -nbr
	}

	for nbr != 0 {

		r = nbr % l
		nbr = nbr / l

		result = append(result, r)
	}

	return result, isNegative
}

func notValid(base string) bool {
	if len(base) < 2 {
		return true
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if (base[i] == base[j]) || base[i] == '-' || base[i] == '+' {
				return true
			}
		}
	}

	return false
}

func NV() {
	printR('N')
	printR('V')
}

func printR(n rune) {
	z01.PrintRune(n)
}
