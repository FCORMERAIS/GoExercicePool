package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var result []int
	counter := 0
	for i := 0; i < len(a)-1; i++ {
		result = append(result, f(a[i], a[i+1]))
	}
	for i := 0; i < len(result); i++ {
		if result[i] < 0 {
			counter++
		}
	}
	if counter == len(result) {
		return true
	}
	counter = 0
	for i := 0; i < len(result); i++ {
		if result[i] > 0 {
			counter++
		}
	}
	if counter == len(result) {
		return true
	}
	return false
}
