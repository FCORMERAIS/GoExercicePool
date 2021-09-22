package piscine

func ToUpper(s string) []rune {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) <= 122 && int(sentence[i]) >= 97 {
			sentence[i] = rune(int(sentence[i]) - 32)
		}
	}
	return sentence
}
