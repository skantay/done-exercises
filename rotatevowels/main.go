package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := getArguments()

	if isValid(arguments) {
		return
	}

	result, toMirror := getMirror(arguments)

	printResult(result, toMirror)
}

func getMirror(arguments []string) (string, []rune) {
	var result string
	var isFirst bool = true
	toMirror := []rune{}

	for _, word := range arguments {
		for _, letter := range word {
			if isVowel(letter) {
				toMirror = append(toMirror, letter)
			}
		}

		if isFirst {
			result = word
			isFirst = false
			continue
		}

		result += " " + word
	}

	return result, toMirror
}

func printResult(result string, toMirror []rune) {
	var current int

	for _, letter := range result {
		if isVowel(letter) {
			z01.PrintRune(toMirror[len(toMirror)-current-1])
			current++
		} else {
			z01.PrintRune(letter)
		}
	}
	z01.PrintRune('\n')
}

func isVowel(letter rune) bool {
	vowels := getVowels()

	for _, v := range vowels {
		if v == letter {
			return true
		}
	}

	return false
}

func getVowels() []rune {
	result := []rune{
		'A',
		'a',
		'E',
		'e',
		'I',
		'i',
		'O',
		'o',
		'U',
		'u',
	}

	return result
}

func getArguments() []string {
	return os.Args[1:]
}

func isValid(arguments []string) bool {
	if len(arguments) < 1 {
		z01.PrintRune('\n')
		return true
	}

	return false
}
