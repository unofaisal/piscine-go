package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	isSign := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			isSign = true
		}
		if c == '-' && !isSign {
			isSign = true
			sign = -1

		}
	}
	return num * sign
}
