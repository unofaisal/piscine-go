package piscine

func Map(f func(int) bool, a []int) []bool {
	bSlice := []bool{}
	for _, in := range a {
		bSlice = append(bSlice, f(in))
	}
	return bSlice
}
