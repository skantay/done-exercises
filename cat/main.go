package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := getArgumetns()
	stdin := true
	for _, fileName := range args {

		if _, err := os.Stat(fileName); err != nil {
			printer("ERROR: open " + fileName + ": no such file or directory\n")
			os.Exit(1)
			return
		}

		content, err := os.ReadFile(fileName)
		if err != nil {
			printer("error")
			os.Exit(1)
		}

		printer(string(content))
		stdin = false
	}

	if stdin {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
		return
	}
}

func printer(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}

func getArgumetns() []string {
	return os.Args[1:]
}
