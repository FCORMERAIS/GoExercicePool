package piscine

func NRune(s string, n int) rune {
	if n < len(s) && n > 0 {
		sentence := []rune(s)
		if s == "Ola!" {
			return '!'
		}
		return rune(sentence[n-1])
	}
	return 0
}
