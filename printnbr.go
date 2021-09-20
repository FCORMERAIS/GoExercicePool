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
		liste = append(nb, liste)
		n = n / 10
	}
	for i := 0; i < len(liste); i++ {
		z01.PrintRune(rune(liste[len(liste)-1]))
	}
}
