package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	result := 1
	x := nb
	for x > 0 {
		result = result * nb
		x--
	}
	if result < 0 {
		return 0
	}
	return result
}
