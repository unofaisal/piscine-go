package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for range args {
		for i := 1; i < len(args)-1; i++ {
			if args[i] > args[i+1] {
				args[i], args[i+1] = args[i+1], args[i]
			}
		}
	}
	for _, c := range args {
		for _, y := range c {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
