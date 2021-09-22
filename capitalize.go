package piscine

func Capitalize(s string) string {
	sentence := []rune(s)
	str := ""
	for i := 0; i < len(sentence); i++ {
		if sentence[i] < 48 || sentence[i] > 57 {
			if sentence[i] < 65 || sentence[i] > 90 {
				if sentence[i] < 97 || sentence[i] > 122 {
					if sentence[i+1] < 97 || sentence[i+1] > 122 {
						sentence[i+1] = sentence[i+1] - 32
					}
				}
			}
		}
		str = str + string(sentence[i])
	}
	return str
}
