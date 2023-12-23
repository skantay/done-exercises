package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n > 9223372036854775807 || n < -9223372036854775807 {

		if n < -9223372036854775807 {
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

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}
	num := n % 10

	zero := '0'
	one := '1'
	two := '2'
	three := '3'
	four := '4'
	five := '5'
	six := '6'
	seven := '7'
	eight := '8'
	nine := '9'

	switch num {
	case 0:
		z01.PrintRune(rune(zero))
	case 1:
		z01.PrintRune(rune(one))
	case 2:
		z01.PrintRune(rune(two))
	case 3:
		z01.PrintRune(rune(three))
	case 4:
		z01.PrintRune(rune(four))
	case 5:
		z01.PrintRune(rune(five))
	case 6:
		z01.PrintRune(rune(six))
	case 7:
		z01.PrintRune(rune(seven))
	case 8:
		z01.PrintRune(rune(eight))
	case 9:
		z01.PrintRune(rune(nine))
	}
}
