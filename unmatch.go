package piscine

func Unmatch(a []int) int {
	mp := make(map[int]int)

	for _, in := range a {
		mp[in] += 1
	}
	n := 0
	for key, val := range mp {
		if val == 1 || val%2 > 0 {
			n = key
		}
	}

	if n != 0 {
		return n
	}
	return -1
}
