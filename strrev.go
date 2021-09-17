package piscine

func StrRev(s string) string {
	result := ""
	letter := ' '
	for i := len(s); i == 0; i-- {
		letter = rune('1')
		result = result + string(letter)
	}
	return result
}
