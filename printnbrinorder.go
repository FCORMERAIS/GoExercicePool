package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	liste := []int{}
	nb := 0
	for k := 0; 0 == n/10; k++ {
		nb = n % 10
		liste = append(liste, nb)
		n = n / 10
	}
	index := 0
	for i := 0; i < len(liste)-1; i++ {
		minimum := 1000000
		for k := i; k < len(liste); k++ {
			if liste[k] < minimum {
				minimum = liste[k]
				index = k
			}
		}
		liste[i], liste[index] = liste[index], liste[i]
	}
	for i := 0; i < len(liste); i++ {
		z01.PrintRune(rune(liste[i]))
	}
}
