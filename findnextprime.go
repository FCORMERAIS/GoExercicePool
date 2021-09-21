package piscine

func FindNextPrime(nb int) int {
	if nb < 1 {
		return 2
	}
	for i := nb; i < 9223372036854775807; i++ {
		pas_bon := 0
		for a := nb / 2; a > 0; a-- {
			if i%a == 0 {
				pas_bon = 1
			}
		}
		if pas_bon == 0 {
			return i
		}
		
	}
	return 0
}
