package piscine

func ToLower(s string) string {
	s1 := []rune(s)
	str := ""
	for i := 0; i <= len(s1)-1; i++ {
		if s1[i] >= 'A' && s1[i] <= 'Z' {
			s1[i] = s1[i] + 32
		}
		str = str + string(s1[i])
	}
	return str
}
