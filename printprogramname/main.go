package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	progName := args[0]

	for i, c := range progName {
		if i > 1 {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
