package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste := make([]int, 1)
	nb := 0
	for n/10 == 0 {
		nb = n % 10
		liste = append(liste, nb)
	}
	for i := 0; i < len(liste); i++ {
		z01.PrintRune(rune(liste[len(liste)-1]))
	}
}
