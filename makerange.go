package piscine

func MakeRange(min, max int) []int {
	size := max - min
	fin2 := make([]int, 0)
	if min > max {
		return []int(nil)
	}
	fin := make([]int, size)
	for i := range fin {
		fin[i] = min + i
	}
	return fin
}
