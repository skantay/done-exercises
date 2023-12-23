package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		word := []rune(v)
		for i := 0; i < len(word); i++ {
			z01.PrintRune(word[i])
		}
		z01.PrintRune('\n')

	}
}
