package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	// Print "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// print points.x
	n := points.x
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		z01.PrintRune('0' + n/10)
		z01.PrintRune('0' + n%10)
	} else {
		z01.PrintRune('0' + n)
	}

	// Print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// print points.y
	m := points.y
	if m < 0 {
		z01.PrintRune('-')
		m = -m
	}
	if m >= 10 {
		z01.PrintRune('0' + m/10)
		z01.PrintRune('0' + m%10)
	} else {
		z01.PrintRune('0' + m)
	}

	// newline
	z01.PrintRune('\n')
}
