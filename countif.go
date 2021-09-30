package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result []bool
	counter := 0
	for i := 0; i < len(a); i++ {
		result = append(result, f(a[i]))
	}
	for i := 0; i < len(result); i++ {
		if result[i] == true {
			counter++
		}
	}
	return counter
}
