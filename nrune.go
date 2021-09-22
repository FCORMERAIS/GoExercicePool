package piscine

func NRune(s string, n int) rune {
	if n < len(s) && n > 0 {
		sentence := []rune(s)
		if len(s) == n {
			return rune(sentence[n-2])
		}
		return rune(sentence[n-1])
	}
	return 0
}
