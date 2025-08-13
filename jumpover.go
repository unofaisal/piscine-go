package piscine

func JumpOver(str string) string {
	s := ""
	i := 2
	for i < len(str) {
		s += string(str[i])
		if !(i+3 < len(str)) {
			break
		}
		i += 3
	}
	return s + "\n"
}