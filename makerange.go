package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	fin := make([]int, size)
	for i := range fin {
		fin[i] = min + i
	}
	return fin
}
