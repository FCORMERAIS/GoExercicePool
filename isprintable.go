package piscine

func IsPrintable(s string) bool {
	counter := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(sentence[i]) <= 126 && int(sentence[i]) >= 32 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	} else {
		return false
	}
}
