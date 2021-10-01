package piscine

func SortWordArr(table []string) {
	for min_index := 0; min_index < len(table)-1; min_index++ {
		for index := min_index + 1; index < len(table); index++ {
			if table[index] < table[min_index] {
				Swap2(&table[index], &table[min_index])
			}
		}
	}
}

func Swap2(a *string, b *string) {
	*a, *b = *b, *a
}
