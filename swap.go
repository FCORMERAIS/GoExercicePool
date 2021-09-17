package piscine

func Swap(a *int, b *int) {
	tempo := *a
	*a = *b
	*b = tempo
}
