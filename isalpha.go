package piscine

func IsAlpha(s string) bool {
	counter := 0
	sentence := []rune(s)
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) <= 90 && int(sentence[i]) >= 65 {
			counter++
		}
		if int(sentence[i]) <= 122 && int(sentence[i]) >= 97 {
			counter++
		}
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
