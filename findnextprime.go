package piscine

func FindNextPrime(nb int) int {
	if nb < 1 {
		return 2
	}
	for i := nb; i < 9223372036854775807; i++ {
		pas_bon := 0
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 || i%11 == 0 || i%13 == 0 {
			pas_bon = 1
		} else {
			for a := nb; a > 0; a-- {
				if i%a == 0 {
					pas_bon = 1
				}
			}
		}
		if pas_bon == 0 {
			return i
		}
	}
	return 0
}
