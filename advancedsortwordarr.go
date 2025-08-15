package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for range a {
		for in, ch := range a {
			if in+1 < len(a) {
				res := f(ch, a[in+1])
				if res == 1 {
					a[in], a[in+1] = a[in+1], a[in]
				}
			}
		}
	}
}
