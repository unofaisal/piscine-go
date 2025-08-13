package piscine

func StringToIntSlice(str string) []int {
	// by := []byte(str)
	in := []int{}
	for _, b := range str {
		in = append(in, int(b))
	}
	return in
}
