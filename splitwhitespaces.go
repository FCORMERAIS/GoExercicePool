package piscine

func SplitWhiteSpaces(s string) []string {
	compteur := 0
	str := ""
	var list1 []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			compteur++
		} else {
			for k := 0; k < compteur; k++ {
				str = str + string(s[i-compteur+k])
			}
			list1 = append(list1, str)
			str = ""
			compteur = 0
		}
	}
	for i := compteur; i > 0; i-- {
		str = str + string(s[len(s)-1-compteur])
	}
	list1 = append(list1, str)
	return list1
}
