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

func ita(in int) string {
	str := ""
	for i := in; i > 0; i /= 10 {
		str = string('0'+(i%10)) + str
	}
	return str
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	printCh("x = ")
	xV := ita(points.x)
	printCh(xV + ", ")
	printCh("y = ")
	yV := ita(points.y)
	printCh(yV)
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
