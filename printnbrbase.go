package piscine

func PrintNbrBase(nbr int, base string) {
	r := []rune(string)
	pos := []int{}

	result := nbr
	outPut := ""
	neg := false
	if nbr < 0 {
		nbr = nbr * -1
		neg = true
	}

	for result >= 0 {
		pos = append(pos, result%len(r)-1)
		result /= len(r)
	}

	for i := len(pos) - 1; i >= 0; i-- {
		outPut += r[i]
	}
	return outPut
}
