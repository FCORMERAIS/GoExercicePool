package piscine

func SortIntegerTable(table []int) {
	minimum := 100000000000
	index := 0
	for i := -1; i < len(table)-1; i++ {
		for k := i; k < len(table)-1; k++ {
			if table[k] < minimum {
				index = k
			}
		}
		table[i], table[index] = table[index], table[i]
	}
}
