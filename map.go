package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for i := 0; i < len(a)-1; i++ {
		result = append(result, f(a[i]))
	}
	return result
}
