package piscine

import (
	"fmt"
	"os"
)

func Inter() {
	args := os.Args[1:]

	if len(args) == 2 {

		w1 := args[0]
		w2 := args[1]

		var result string

		word1 := []rune(w1)
		word2 := []rune(w2)

		for i := 0; i < len(word1); i++ {
			for j := i; j < len(word2); j++ {
				if word1[i] == word2[j] && doesContainDouble(word1[i], result) {
					result += string(word1[i])
					break
				}
			}
		}

		fmt.Println(result)
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
