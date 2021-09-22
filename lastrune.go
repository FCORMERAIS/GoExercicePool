package piscine

func LastRune(s string) rune {
	sentence := []rune(s)
	return rune(sentence[len(sentence)-1])
}
