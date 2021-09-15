package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	rune1 := '0'
	rune2 := '0'
	rune3 := '0'
	rune4 := '0'
	for i := 0; i < 100; i++ {
		for k := 0; k < 100; k++ {
			rune4++
			z01.PrintRune(rune(rune1))
			z01.PrintRune(rune(rune2))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune(rune3))
			z01.PrintRune(rune(rune4))
			if rune4 == 9 {
				rune4 = 0
				rune3++
			}
			if rune3 == 9{
				rune2++
				rune3 = 0
			}
			if rune2 == 9{
				rune1++
				rune2 = 0
			}
		}
	}
}