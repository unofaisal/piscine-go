package piscine

// import "fmt"

func Atoi(s string) int {
	var f int
	sign := 1
	if !(len(s) == 0) {

		if s[0] == '-' {
			sign = -1
			s = s[1:]
		} else if s[0] == '+' {
			sign = 1
			s = s[1:]
		}

		for _, v := range s {
			if v < '0' || v > '9' {
				// fmt.Println(int(v - '0'))
				return 0
			}
			// fmt.Println((v - '0'))
			f = f*10 + int(v-'0')
		}
	}
	return f * sign
}
