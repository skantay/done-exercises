package piscine

import "github.com/01-edu/z01"

func PrintNbrTest(n int) {
	chars := make(map[int]int)
	chars[0] = 48
	chars[1] = 49
	chars[2] = 50
	chars[3] = 51
	chars[4] = 52
	chars[5] = 53
	chars[6] = 54
	chars[7] = 55
	chars[8] = 56
	chars[9] = 57

	if n <= 9 && n >= 0 {
		z01.PrintRune(rune(chars[n]))
		return
	}

	var isNegative bool = false
	var reversed int
	var len int = 0
	hasNonZeroDigit := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n != 0 {

		digit := n % 10

		if digit != 0 {
			hasNonZeroDigit = true
			reversed = reversed*10 + digit
		} else if hasNonZeroDigit {
			reversed = reversed * 10
		}

		n = n / 10
		len++
	}

	for i := len; i > 0; i-- {
		digit := 0

		digit = reversed % 10
		reversed = reversed / 10
		if isNegative {
			z01.PrintRune('-')

			isNegative = false
		}
		z01.PrintRune(rune(chars[digit]))
	}

	return
}
