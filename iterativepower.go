package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 50 {
		return 0
	}
	return nb ^ power
}
