package piscine

func SortIntegerTable(table []int) {
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
}
