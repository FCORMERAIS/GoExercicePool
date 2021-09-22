package piscine

func NRune(s string, n int) rune {
	if n < len(s) && n > 0 {
		sentence := []rune(s)
		return rune(sentence[n-2])
	}
	return 0
}
