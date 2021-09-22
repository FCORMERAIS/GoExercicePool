package piscine

func IsUpper(s string) bool {
	counter := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(sentence[i]) <= 90 && int(sentence[i]) >= 65 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	} else {
		return false
	}
}
