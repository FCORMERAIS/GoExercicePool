package piscine

func ToUpper(s string) string {
	sentence := []rune(s)
	new_sentence := ""
	for i := 0; i <= len(sentence)-1; i++ {
		if sentence[i] <= 'z' && sentence[i] >= 'a' {
			sentence[i] = sentence[i] - 32
		}
		new_sentence = sentence + string(sentence[i])
	}
	return new_sentence
}
