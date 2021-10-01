package piscine

import (
	"os"
)

func SortIntegerTabl(table []string) []string {
	index := 0
	if len(table) == 1 {
		return table
	}
	for i := 0; i <= len(table)-1; i++ {
		minimum := table[i]
		for k := i; k < len(table); k++ {
			if table[k] <= minimum {
				minimum = table[k]
				index = k
			}
		}
		table[i], table[index] = table[index], table[i]
	}
	return table
}

func SortWordArr(a []string) {
	sentence := SortIntegerTabl(a)
	for i := 0; i < len(sentence)-1; i++ {
		os.Stderr.WriteString(sentence[i])
	}
	os.Stderr.WriteString("\n")
}
