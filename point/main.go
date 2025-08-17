package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printCh(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func printInt(n int) {
	first := n / 10
	second := n % 10

	digit1 := '0'
	for i := 0; i < first; i++ {
		digit1++
	}
	z01.PrintRune(digit1)

	digit2 := '0'
	for i := 0; i < second; i++ {
		digit2++
	}
	z01.PrintRune(digit2)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	printCh("x = ")
	printInt(points.x)
	printCh(", y = ")
	printInt(points.y)
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
