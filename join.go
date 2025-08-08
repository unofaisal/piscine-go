package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i, w := range strs {
		s += w
		if i != len(strs)-1 {
			s += sep
		}
	}
	return s
}
