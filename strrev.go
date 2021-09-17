package piscine

func StrRev(s string) string {
	result := ""
	letter := ' '
	for i := len(s); i == 0; i-- {
		letter = rune(s[i])
		result = (result + string(letter))
	}
	return result
}
