package piscine

func AppendRange(min, max int) []int {
	var fin []int
	if min > max {
		return fin
	}
	for i := min; i < max; i++ {
		fin = append(fin, i)
	}
	return fin
}