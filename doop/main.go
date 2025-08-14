package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	opp := os.Args[2]
	b, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
		return // invalid number, no output
	}

	switch opp {
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
	}

	if !checkOverFlow(a, b, opp) {
		return // overflow or invalid op, no output
	}

	switch opp {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "%":
		fmt.Println(a % b)
	}
}

func checkOverFlow(a, b int64, opp string) bool {
	const MaxInt64 = 9223372036854775807
	const MinInt64 = -9223372036854775808

	switch opp {
	case "+":
		if b > 0 && a > MaxInt64-b {
			return false
		}
		if b < 0 && a < MinInt64-b {
			return false
		}
		return true
	case "-":
		if b < 0 && a > MaxInt64+b {
			return false
		}
		if b > 0 && a < MinInt64+b {
			return false
		}
		return true
	case "*":
		if a == 0 || b == 0 {
			return true
		}
		if (a == -1 && b == MinInt64) || (b == -1 && a == MinInt64) {
			return false
		}
		if a > 0 {
			if b > 0 && a > MaxInt64/b {
				return false
			}
			if b < 0 && b < MinInt64/a {
				return false
			}
		} else {
			if b > 0 && a < MinInt64/b {
				return false
			}
			if b < 0 && a < MaxInt64/b {
				return false
			}
		}
		return true
	case "/":
		if b == 0 {
			return false
		}
		if a == MinInt64 && b == -1 {
			return false
		}
		return true
	case "%":
		if b == 0 {
			return false
		}
		return true
	default:
		return false
	}
}
