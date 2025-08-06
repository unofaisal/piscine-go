package piscine

func IsAlpha(s string) bool {
	x := true
	for _, c := range s {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
			x = false
		}
	}
	return x
}
