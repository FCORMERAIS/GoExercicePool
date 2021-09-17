package piscine

func UltimateDivMod(a *int, b *int) {
	tempo := &a
	*a = *a / *b
	*b = tempo % *b
}
