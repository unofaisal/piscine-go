package piscine

func SortWordArr(a []string) {
	for range a {
		for ind, str := range a {
			if ind < len(a)-1 && str > a[ind+1] {
				a[ind], a[ind+1] = a[ind+1], a[ind]
			}
		}
	}
}
