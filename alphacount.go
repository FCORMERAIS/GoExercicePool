package piscine

func AlphaCount(s string) int {
	counter := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(sentence[i]) <= 90 && int(sentence[i]) >= 65 {
			counter++
		}
		if int(sentence[i]) <= 122 && int(sentence[i]) >= 97 {
			counter++
		}
	}
	return counter
}
