package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for nb > 0 {
		result = result * nb
		nb--
	}
	return result
}
