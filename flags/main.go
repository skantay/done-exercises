package main

import (
	"fmt"
	"os"
)

func main() {
	// get arguments
	// find flags
	// flags functions
	// print elemenets considering flags

	var textHelp string = `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.`

	arguments := getArguments()
	isHelp, isInsert, toInsert, isOrder := getFlags(arguments)

	if isHelp || len(arguments) == 0 {
		fmt.Println(textHelp)
		return
	}

	result := createResult(arguments)

	if isInsert {
		result = result + toInsert
	}

	if isOrder {
		result = order(result)
	}

	fmt.Println(result)
}

func insert(result, toInsert string) string {
	var clean string
	var t bool
	toInsertRunes := []rune(toInsert)

	for i := 0; i < len(toInsertRunes); i++ {
		if toInsertRunes[i] == '=' && !t {
			t = true
			continue
		}

		if t {
			clean += string(toInsertRunes[i])
		}
	}

	return result + clean
}

func createResult(arguments []string) string {
	var result string
	for _, word := range arguments {
		if word != "FLAG FOUND" {
			result = result + word
		}
	}
	return result
}

func order(result string) string {
	letters := []rune(result)

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			a := letters[j]
			b := letters[j+1]
			if a == b {
				continue
			}
			if a > b {
				letters[j], letters[j+1] = letters[j+1], letters[j]
			}
		}
	}

	return string(letters)
}

func getFlags(arguments []string) (bool, bool, string, bool) {
	var isHelp bool = false
	var isOrder bool = false
	var isInsert bool = false
	var toInsert string

	for i, flag := range arguments {
		if flag == "--help" || flag == "-h" {
			isHelp = true
		} else if flag == "--order" || flag == "-o" {
			isOrder = true
			arguments[i] = "FLAG FOUND"
		} else if firstN(flag, 2) == "-i" || firstN(flag, 8) == "--insert" {
			isInsert = true
			if firstN(flag, 8) == "--insert" {
				toInsert = flag[9:]
				arguments[i] = "FLAG FOUND"
				continue
			}
			toInsert = flag[3:]
			arguments[i] = "FLAG FOUND"
		}
	}

	return isHelp, isInsert, toInsert, isOrder
}

func firstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func getArguments() []string {
	return os.Args[1:]
}
