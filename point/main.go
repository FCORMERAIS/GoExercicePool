package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x1 rune
	x2 rune
	y1 rune
	y2 rune
}

func setPoint(ptr *point) {
	ptr.x1 = '4'
	ptr.x2 = '2'
	ptr.y1 = '2'
	ptr.y2 = '1'
}

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune(points.x1))
	z01.PrintRune(rune(points.x2))
	z01.PrintRune(rune(','))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune(points.y1))
	z01.PrintRune(rune(points.y2))
}
