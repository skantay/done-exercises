package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		word1 := []rune(args[0])
		word2 := []rune(args[1])

		var result string
		var index int

		for i := 0; i < len(word1); i++ {
			for j := index; j < len(word2); j++ {
				for word1[i] == word2[j] && doesContainDouble(rune(word1[i]), result) {
					result += string(word1[i])
					index = j
					break
				}
			}
		}

		if result == string(word1) {
			z01.PrintRune('1')
			z01.PrintRune('\n')
			return
		}
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}

func doesContainDouble(word rune, result string) bool {
	for _, v := range result {
		if word == v {
			return false
		}
	}
	return true
}
