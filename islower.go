package piscine

func IsLower(s string) bool {
	counter := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if int(sentence[i]) <= 122 && int(sentence[i]) >= 97 {
			counter++
		}
	}
	if counter == len(s) {
		return true
	} else {
		return false
	}
}
