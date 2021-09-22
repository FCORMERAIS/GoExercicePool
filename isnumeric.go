package piscine

func IsNumeric(s string) bool {
	counter := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(sentence[i]) <= 57 && int(sentence[i]) >= 48 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	} else {
		return false
	}
}
