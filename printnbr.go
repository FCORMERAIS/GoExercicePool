package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste := []int{}
	nb := 0
	if n < 0 {
		z01.PrintRune(rune('-'))
	}
	for k := 0; 0 == n/10; k++ {
		nb = n % 10
		liste = append(liste, nb)
		n = n / 10
	}
	for i := 0; i > len(liste)-1; i++ {
		z01.PrintRune(rune(liste[i]))
	}
}
