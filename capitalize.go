package piscine

func Capitalize(s string) string {
	r := []rune(s)

	for i, c := range r {
		if c >= 'A' && c <= 'Z' {
			r[i] = r[i] - 'A' + 'a'
		}
	}

	for i, c := range r {
		if i == 0 {
			if r[0] >= 'a' && r[0] <= 'z' {
				r[0] = 'A' + r[0] - 'a'
			}
		}
		if i < len(r)-1 && !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
			if r[i+1] >= 'a' && r[i+1] <= 'z' {
				r[i+1] = 'A' + r[i+1] - 'a'
			}
		}
	}
	return string(r)
}
