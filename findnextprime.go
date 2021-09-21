package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	tmp := 0
	for i := 0; (nb+i%2 != 0) && (nb+i%3 != 0) && (nb+i%5 != 0); i++ {
		tmp = i
	}
	return nb + tmp
}
