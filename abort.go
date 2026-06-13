package piscine
// 
func Abort(a, b, c, d, e int) int {
	sl := []int{a, b, c, d, e}
	for range sl {
		for i, in := range sl {
			if i+1 < len(sl) && in > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
	return sl[len(sl)/2]
}
