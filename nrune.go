package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	var x rune
	for i, c := range r {
		if i == n-1 {
			x = c
		}
	}
	return x
}
