package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	args[0] = args[len(args)-1]
	args = args[:len(args)-1]

	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			a := []rune(args[j])
			b := []rune(args[j+1])
			if a[0] == b[0] {
				continue
			}
			if a[0] > b[0] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
	for i := 0; i < len(args); i++ {
		word := []rune(args[i])
		for j := 0; j < len(word); j++ {
			z01.PrintRune(word[j])
		}
		z01.PrintRune('\n')
	}
}
