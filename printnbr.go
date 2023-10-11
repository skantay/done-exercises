package piscine 

import "github.com/01-edu/z01"

func Printnbr(n int) {
	var result int
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	for n != 0 {
		result = (result * 10) + (n % 10)
		n /= 10
	}

	for result != 0 {
		z01.PrintRune(rune(result % 10)+'0')
		result /= 10
	}
}