package piscine

func Split(s, sep string) []string {
	compteur := 0
	str := ""
	var list1 []string
	for i := 0; i < len(s); i++ {
		if s[i] != sep[0] {
			compteur++
		} else if s[i] == sep[0] && compteur != 0 {
			compteur_sep := 0
			for j := 0; j < len(sep); j++ {
				if len(s)-i >= len(sep) && s[i+j] == sep[j] {
					compteur_sep++
				}
			}
			if compteur_sep == len(sep) {
				for k := 0; k < compteur; k++ {
					if i-compteur+k == -1 {
						return list1
					}
					str = str + string(s[i-compteur+k])
				}
				list1 = append(list1, str)
				str = ""
				compteur = 0
				i = i + len(sep) - 1
			} else {
				compteur++
			}
		}
	}
	for i := compteur; i > 0; i-- {
		compteur--
		str = str + string(s[len(s)-1-compteur])
	}
	list1 = append(list1, str)
	return list1
}
