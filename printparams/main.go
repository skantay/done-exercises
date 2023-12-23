package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var word string
	for i := 1; i < len(args); i++ {
		word = args[i]
		letters := []rune(word)
		for j := 0; j < len(args[i]); j++ {
			z01.PrintRune(letters[j])
		}
		z01.PrintRune('\n')
	}
}
