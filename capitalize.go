package piscine

func Capitalize(s string) string {
	sentence := []rune(s)
	str := ""
	if sentence[0] >= 'a' && sentence[0] <= 'z' {
		sentence[0] = sentence[0] - 32
	}
	str = str + string(sentence[0])
	for i := 1; i < len(sentence); i++ {
		if sentence[i] >= 'a' && sentence[i] <= 'z' {
			if sentence[i-1] >= 'a' && sentence[i-1] <= 'z' || sentence[i-1] >= '0' && sentence[i-1] <= '9' || sentence[i-1] >= 'A' && sentence[i-1] <= 'Z' {
				sentence[i] = sentence[i] + 0
			} else {
				sentence[i] = sentence[i] - 32
			}
		}
		if sentence[i] >= 'A' && sentence[i] <= 'Z' {
			if sentence[i-1] >= 'a' && sentence[i-1] <= 'z' || sentence[i-1] >= '0' && sentence[i-1] <= '9' || sentence[i-1] >= 'A' && sentence[i-1] <= 'Z' {
				sentence[i] = sentence[i] + 32
			} else {
				sentence[i] = sentence[i] + 0
			}
		}
		str = str + string(sentence[i])
	}
	return str
}
