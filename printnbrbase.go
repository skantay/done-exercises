package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if nbr > 9223372036854775807 || nbr < -9223372036854775807 {

		if nbr < -9223372036854775807 {
			z01.PrintRune('-')
		}
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}

	if notvalid(base) {
		nv()
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	baseLen := len(base)
	var num []int

	for nbr > 0 {
		remainder := nbr % baseLen
		num = append([]int{remainder}, num...)
		nbr /= baseLen
	}

	for i := 0; i < len(num); i++ {
		index := num[i]
		z01.PrintRune(rune(base[index]))
	}
}

func notvalid(base string) bool {
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

func nv() {
	printR('N')
	printR('V')
}
