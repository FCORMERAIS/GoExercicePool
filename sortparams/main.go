package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable(table []string) []string {
	index := 0
	if len(table) == 1 {
		return table
	}
	for i := 0; i < len(table)-1; i++ {
		minimum := table[0]
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

func main() {
	argument := os.Args
	list := []string{}
	for i := 1; i < len(argument); i++ {
		list = append(list, argument[i])
	}
	list = SortIntegerTable(list)
	for i := 0; i < len(argument)-1; i++ {
		for k := 0; k < len(list[i])-1; k++ {
			z01.PrintRune(rune(list[i][k]))
		}
		z01.PrintRune(rune('\n'))
	}
}
