package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	k := []rune(s)
	for _, char := range k {
		z01.PrintRune(char)
	}
}
