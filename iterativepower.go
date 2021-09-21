package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 50 {
		return 0
	}
	for i := 0; i < power; i++ {
		nb = nb * power
	}
	return nb
}
