package piscine

func ToUpper(s string) []rune {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) <= 'z' && int(sentence[i]) >= 'a' {
			sentence[i] = rune(int(sentence[i]) - 32)
		}
	}
	return sentence
}
