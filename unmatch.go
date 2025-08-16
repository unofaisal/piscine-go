package piscine

func Unmatch(a []int) int {
	mpC := make(map[int]int)

	for _, in := range a {
		mpC[in] += 1
	}
	for key, val := range mpC {
		if val == 1 || val%2 > 0 {
			return key
		}
	}
	return -1
}
