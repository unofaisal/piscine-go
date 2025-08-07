package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i, j := range args {
		if i != 0 {
			for _, y := range j {
				z01.PrintRune(y)
			}
			z01.PrintRune('\n')
		}
	}
}
