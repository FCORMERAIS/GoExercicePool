package piscine

func Split(s, sep string) []string {
	var list1 []string
	word := ""
	testsep := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) < len(s) {
			for k := i; k < i+len(sep); k++ {
				testsep = testsep + string(s[k])
			}
		}
		if i == len(s)-1 {
			word = word + string(s[i])
			list1 = append(list1, word)
			word = ""
		} else if testsep == sep {
			list1 = append(list1, word)
			word = ""
			i = i + len(sep) - 1
		} else {
			if string(s[i]) == " " {
			} else {
				word = word + string(s[i])
			}
		}
		testsep = ""
	}
	return list1
}
