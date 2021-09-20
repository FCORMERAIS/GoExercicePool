package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste := []int{}
	nb := 0
	if n < 0 {
		z01.PrintRune(rune('-'))
	}
	for n/10 == 0 {
		nb = n % 10
		liste = append(liste, nb)
		n = n / 10
	}
	for i := len(liste) - 1; i > 0; i-- {
		z01.PrintRune(rune(liste[i]))
	}
}
