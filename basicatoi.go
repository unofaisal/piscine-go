package piscine

func BasicAtoi(s string) int {
	i := 0
	for _, char := range s {
		x := int(char - '0')

		i = i*10 + x
	}
	return i
}
