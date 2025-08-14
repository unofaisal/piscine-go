package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asce := true
	desc := true
	for ind, in := range a {
		if ind < len(a)-1 {
			sortR := f(in, a[ind+1])
			if sortR < 0 {
				desc = false
			}
			if sortR > 0 {
				asce = false
			}
		}
	}
	return asce || desc
}

// func gtLt(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else if a < b {
// 		return -1
// 	} else {
// 		return 0
// 	}
// }
