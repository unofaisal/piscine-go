package main

import (
	"os"
)

func isNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, r := range s {
		if i == 0 && r == '-' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func atoi(s string) (int64, bool) {
	if !isNumeric(s) || len(s) > 19 {
		return 0, false
	}
	neg := false
	start := 0
	if s[0] == '-' {
		neg = true
		start = 1
	}
	var result int64
	for i := start; i < len(s); i++ {
		d := int64(s[i] - '0')
		if result > (9223372036854775807-d)/10 {
			return 0, false
		}
		result = result*10 + d
	}
	if neg {
		result = -result
	}
	return result, true
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func printNum(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	buf := []byte{}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	for n > 0 {
		buf = append([]byte{byte(n%10) + '0'}, buf...)
		n /= 10
	}
	if neg {
		buf = append([]byte{'-'}, buf...)
	}
	buf = append(buf, '\n')
	os.Stdout.Write(buf)
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	op := os.Args[2]
	arg2 := os.Args[3]

	if len(op) != 1 || (op != "+" && op != "-" && op != "*" && op != "/" && op != "%") {
		return
	}

	num1, ok1 := atoi(arg1)
	num2, ok2 := atoi(arg2)

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		if num1 > 0 && num2 > 0 && num1 > (9223372036854775807-num2) ||
			num1 < 0 && num2 < 0 && num1 < (-9223372036854775808-num2) {
			return
		}
		printNum(num1 + num2)
	case "-":
		if num1 > 0 && num2 < 0 && num1 > (9223372036854775807+num2) ||
			num1 < 0 && num2 > 0 && num1 < (-9223372036854775808+num2) {
			return
		}
		printNum(num1 - num2)
	case "*":
		if num1 != 0 && num2 != 0 &&
			(num1 > 9223372036854775807/num2 ||
				num1 < -9223372036854775808/num2) {
			return
		}
		printNum(num1 * num2)
	case "/":
		if num2 == 0 {
			printStr("No division by 0")
			return
		}
		printNum(num1 / num2)
	case "%":
		if num2 == 0 {
			printStr("No modulo by 0")
			return
		}
		printNum(num1 % num2)
	}
}
