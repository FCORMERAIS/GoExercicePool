package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	if nb == 3 {
		return 3
	}
	if nb == 4 || nb == 5 {
		return 5
	}
	if nb == 6 || nb == 7 {
		return 7
	}
	if nb == 8 || nb == 9 || nb == 10 || nb == 11 {
		return 11
	}
	if nb == 12 || nb == 13 {
		return 13
	}
	for i := nb; i < 9223372036854775807; i++ {
		pas_bon := 0
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 || i%11 == 0 || i%13 == 0 {
			pas_bon = 1
		} else {
			for a := 13; a < i; a++ {
				if i%a == 0 {
					pas_bon = 1
				}
				a++
			}
		}
		if pas_bon == 0 {
			return i
		}
	}
	return 0
}
