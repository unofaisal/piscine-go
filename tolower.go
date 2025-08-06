package piscine

func ToLower(s string) string {
	final := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			k := 'a' + (c - 'A')
			final += string(k)
		} else {
			final += string(c)
		}
	}
	return final
}
