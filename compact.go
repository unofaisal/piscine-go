package piscine

func Compact(ptr *[]string) int {
	nS := []string{}
	for _, str := range *ptr {
		if str != "" {
			nS = append(nS, str)
		}
	}
	*ptr = append((*ptr)[:0], nS...)
	return len(nS)
}
