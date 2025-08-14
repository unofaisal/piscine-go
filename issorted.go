package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isSort := false
	for ind, in := range a {
		if ind < len(a)-1 {
			if f(in, a[ind+1]) == 1 {
				return false
			} else if f(in, a[ind+1]) == 0 {
				isSort = true
			} else {
				isSort = true
			}
		}
	}
	return isSort
}

// func gtLt(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else if a < b {
// 		return 0
// 	} else {
// 		return -1
// 	}
// }
