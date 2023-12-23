package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println(alphamirror(args[0]))
}

func alphamirror(word string) string {
	reversedLower := createRevAlpLower()
	lower := createAlpLower()
	reversedUpper := createRevAlpUpper()
	upper := createAlpUpper()

	var result string

	for _, v := range word {
		if v >= 'a' && v <= 'z' {
			index := getIndex(lower, v)
			result += string(reversedLower[index])
			continue
		} else if v >= 'A' && v <= 'Z' {
			index := getIndex(upper, v)
			result += string(reversedUpper[index])
			continue
		}
		result += string(v)
	}
	return result
}

func getIndex(alphabet []rune, letter rune) int {
	for i, v := range alphabet {
		if v == letter {
			return i
		}
	}
	return 0
}

func createRevAlpLower() []rune {
	var result []rune

	char := 'z'

	for ; char >= 'a'; char-- {
		result = append(result, char)
	}

	return result
}

func createAlpLower() []rune {
	var result []rune

	char := 'a'

	for ; char <= 'z'; char++ {
		result = append(result, char)
	}

	return result
}

func createRevAlpUpper() []rune {
	var result []rune

	char := 'Z'

	for ; char >= 'A'; char-- {
		result = append(result, char)
	}

	return result
}

func createAlpUpper() []rune {
	var result []rune

	char := 'A'

	for ; char <= 'Z'; char++ {
		result = append(result, char)
	}

	return result
}
