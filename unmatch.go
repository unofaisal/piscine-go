package piscine

func Unmatch(a []int) int {
	mp := make(map[int]int)

	for _, in := range a {
		mp[in] += 1
	}

	// for key, val := range mp {
	// 	if val == 1 || val%2 > 0 {
	// 		return key
	// 	}
	// }
	for _, i := range a {
		for key, ind := range mp {
			if key == i && ind%2 != 0 {
				return i
			}
		}
	}

	return -1
}
