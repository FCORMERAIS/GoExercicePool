package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 50 {
		return 0
	}
	result := nb
	for i := 0; i < power-1; i++ {
		result = result * nb
	}
	return result
}
