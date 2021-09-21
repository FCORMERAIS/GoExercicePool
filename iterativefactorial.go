package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 10 {
		return 0
	}
	result := 1
	for i := 1; i < nb; i++ {
		result = result + i
	}
	return result
}
