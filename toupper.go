package piscine

func ToUpper(s string) string {
	final := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			k := 'A' + (c - 'a')
			final += string(k)
		} else {
			final += string(c)
		}
	}
	return final
}
