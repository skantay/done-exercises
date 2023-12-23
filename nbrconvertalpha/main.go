package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	args = args[1:]
	if len(args) == 0 {
		return
	}
	alphabet := createAlphabet(flag(args))

	for i := 0; i < len(args); i++ {
		letter := args[i]

		if letter == "--upper" {
			continue
		}

		digit := letterToInt(letter)

		if digit < 27 {
			z01.PrintRune(alphabet[digit])
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func flag(args []string) bool {
	if len(args) == 0 {
		return false
	}

	for i := 0; i < len(args); i++ {
		if args[i][0] == '-' && args[i][1] == '-' {
			return true
		}
	}
	return false
}

func createAlphabet(isUpper bool) [27]rune {
	var alphabet [27]rune
	var toUpper rune
	letter := 'a'
	if isUpper {
		toUpper = 32
	}
	for i := 1; i < 27; i++ {
		alphabet[i] = letter - rune(toUpper)
		letter++
	}

	return alphabet
}

func letterToInt(letter string) int {
	var result int

	for i, v := range letter {

		digit := int(v - '0')

		result += digit

		if i != len(letter)-1 {
			result = result * 10
		}
	}

	return result
}
