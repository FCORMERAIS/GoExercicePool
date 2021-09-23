package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable(table []int) []int {
	index := 0
	for i := 0; i < len(table)-1; i++ {
		minimum := 1000000
		for k := i; k < len(table); k++ {
			if table[k] < minimum {
				minimum = table[k]
				index = k
			}
		}
		table[i], table[index] = table[index], table[i]
	}
	return table
}

func FirstRune(s string) rune {
	sentence := []rune(s)
	return rune(sentence[0])
}

func main() {
	argument := os.Args
	list := []int{}
	for i := 1; i < len(argument); i++ {
		list = append(list, int(FirstRune(argument[i])))
	}
	list = SortIntegerTable(list)
	for i := 1; i < len(argument); i++ {
		z01.PrintRune(rune(list[i]))
		z01.PrintRune(rune('\n'))
	}
}
