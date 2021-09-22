package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	return rune(sentence[n-1])
}
