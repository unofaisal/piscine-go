package piscine

func ActiveBits(n int) int {
	r := []int{}

	for i := n; i > 0; i /= 2 {
		r = append(r, i%2)
	}
	counter := 0
	for _, i := range r {
		if i == 1 {
			counter++
		}
	}
	return counter
}
