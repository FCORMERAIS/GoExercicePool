package piscine

func ToUpper(s string) []rune {
	sentence := []rune(s)
	new_sentence := ""
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) <= 122 && int(sentence[i]) >= 97 {
			sentence[i] = rune(int(sentence[i]) - 32)
		}
		new_sentence = sentence + string(sentence[i])
	}
	return new_sentence
}
