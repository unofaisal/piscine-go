package main

import "github.com/01-edu/z01"

func main() {
	i := 48

	for i <= (48 + 9) {
		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune('\n')
}
