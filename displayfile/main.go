package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	args = args[1:]

	if len(args) < 1 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	dat, _ := os.ReadFile(args[0])
	fmt.Print(string(dat))
}
