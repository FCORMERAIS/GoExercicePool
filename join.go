package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		str = str + string(strs[i])
		if i+1 != len(strs) {
			str = str + sep
		}
	}
	return str
}
