package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	s := []rune{}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		x := n % 10
		s = append(s, rune(x)+'0')
		n /= 10

	}
	for idx, nb := range s {
		if idx < len(s)-1 && nb > s[idx+1] {
			s[idx], s[idx+1] = s[idx+1], s[idx]
		}
	}
	for _, c := range s {
		z01.PrintRune(c)
	}
}
