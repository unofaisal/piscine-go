package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	sl := []string{}
	for i, l := range s {

		if l >= ' ' && l <= '~' {
			word += string(l)
		}
		if word[len(word)-1] == ' ' || word[len(word)-1] == '	' || word[len(word)-1] == '\n' {
			sl = append(sl, word[:len(word)-1])
			word = ""

		}
		if i == len(s)-1 {
			sl = append(sl, word[:len(word)])
		}
	}
	return sl
}
