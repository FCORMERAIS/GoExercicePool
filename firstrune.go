package piscine

func FirstRune(s string) rune {
	sentence := []byte(s)
	return rune(sentence[0])
}
