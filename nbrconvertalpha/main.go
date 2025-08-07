package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, c := range args {
		for _, y := range c {
			z01.PrintRune(rune('a' + int(y)))
		}
	}
}
