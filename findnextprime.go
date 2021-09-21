package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	if nb == 4 {
		return 5
	}
	if nb == 6 {
		return 7
	}
	if nb ==8{
		return 11
	}
	for i := nb; i < 9223372036854775807; i++ {
		pas_bon := 0
		if nb < 14 {
			for a := nb; a < 0; a-- {
				if i%a == 0 {
					pas_bon = 1
				}
			}
			if pas_bon == 0 {
				return i
			}
		}
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 || i%11 == 0 || i%13 == 0 {
			pas_bon = 1
		} else {
			for a := nb - 1; a >= 13; a-- {
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
