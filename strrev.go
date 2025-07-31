package piscine

func StrRev(s string) string {
	r := []rune(s)
	var x string
	for i := len(r) - 1; i >= 0; i-- {
		x += string(r[i])
	}
	return x
}
