package piscine

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		sentence := []rune(s)
		return rune(sentence[n-1])
	}
	return 0
}
