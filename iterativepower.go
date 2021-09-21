package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 50 {
		return 0
	}
	result := nb
	for i := 0; i < power; i++ {
		result = result * power
	}
	return result
}
