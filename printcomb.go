package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	chiffre1 := '0'
	chiffre2 := '0'
	chiffre3 := '0'
	for i := 0; i < 600; i++ {
		chiffre1++
		if chiffre1 != chiffre2 && chiffre2 != chiffre3 && chiffre3 != chiffre1 && chiffre3 < chiffre2 && chiffre2 < chiffre1 && chiffre3 != 7 {
			z01.PrintRune(rune(chiffre3))
			z01.PrintRune(rune(chiffre2))
			z01.PrintRune(rune(chiffre1))
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		if chiffre1 == '9' {
			chiffre1 = '0'
			chiffre2++
		}
		if chiffre2 == '9' {
			chiffre3++
			chiffre2 = '0'
		}
		if i == 699 {
			z01.PrintRune(rune('7'))
			z01.PrintRune(rune('8'))
			z01.PrintRune(rune('9'))
			z01.PrintRune('\n')
		}

	}
}
