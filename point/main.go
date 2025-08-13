package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		return "-" + itoa(-n)
	}
	digits := ""
	for n > 0 {
		d := n % 10
		digits = string('0'+d) + digits
		n /= 10
	}
	return digits
}

func main() {
	points := &point{}

	setPoint(points)
	printStr("x = ")
	printStr(itoa(points.x))
	printStr(", y = ")
	printStr(itoa(points.y))
	printStr("\n")
}
