package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args); i > 0; i-- {
		for _, y := range args[i] {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
