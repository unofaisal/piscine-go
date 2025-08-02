package piscine

func BasicAtoi2(s string) int {
	var x int
	for _, i := range s {
		if !(i >= '0' && i <= '9') {
			return 0
		} else {
			x = x*10 + int(i-'0')
		}
	}
	return x
}
