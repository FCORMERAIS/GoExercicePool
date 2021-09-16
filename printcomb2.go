package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	rune1 := '0'
	rune2 := '0'
	rune3 := '0'
	rune4 := '0'
	for rune1 != '9' || rune2 != '8' {
		rune4++
		if rune1 < rune3 {
			z01.PrintRune(rune(rune1))
			z01.PrintRune(rune(rune2))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune(rune3))
			z01.PrintRune(rune(rune4))
			z01.PrintRune(rune(','))
			z01.PrintRune(rune(' '))
		} else if rune1 == rune3 {
			if rune2 < rune4 {
				z01.PrintRune(rune(rune1))
				z01.PrintRune(rune(rune2))
				z01.PrintRune(rune(' '))
				z01.PrintRune(rune(rune3))
				z01.PrintRune(rune(rune4))
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
		if rune2 == '9' && rune3 == '9' && rune4 == '9' {
			rune1++
			rune2 = '0'
			rune3 = '0'
			rune4 = '0'
			z01.PrintRune(rune(rune1))
			z01.PrintRune(rune(rune2))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune(rune3))
			z01.PrintRune(rune(rune4))
			z01.PrintRune(rune(','))
			z01.PrintRune(rune(' '))
		}
		if rune3 == '9' && rune4 == '9' {
			rune2++
			rune3 = '0'
			rune4 = '0'
			z01.PrintRune(rune(rune1))
			z01.PrintRune(rune(rune2))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune(rune3))
			z01.PrintRune(rune(rune4))
			z01.PrintRune(rune(','))
			z01.PrintRune(rune(' '))
		}
		if rune4 == '9' {
			rune4 = '0'
			rune3++
			z01.PrintRune(rune(rune1))
			z01.PrintRune(rune(rune2))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune(rune3))
			z01.PrintRune(rune(rune4))
			z01.PrintRune(rune(','))
			z01.PrintRune(rune(' '))
		}
	}
	z01.PrintRune(rune(rune1))
	z01.PrintRune(rune(rune2))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('9'))
	z01.PrintRune(rune('9'))
	z01.PrintRune(rune('\n'))
}
