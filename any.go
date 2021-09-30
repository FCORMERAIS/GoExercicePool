
package piscine

func Any(f func(string) bool, a []string) bool {
	var result []bool
	for i := 0; i < len(a); i++ {
		result = append(result, f(a[i]))
	}
	for i := 0; i < len(result); i++ {
		if result[i] == true {
			return true
		}
	}
	return false
}
