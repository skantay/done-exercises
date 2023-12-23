package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argu := os.Args
	word := []rune(argu[0])
	for i := 2; i < len(word); i++ {
		z01.PrintRune(word[i])
	}
	z01.PrintRune('\n')
}
