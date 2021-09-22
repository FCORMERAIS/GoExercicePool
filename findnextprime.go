package piscine

func FindNextPrime(nb int) int {
	c := 0
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
	if nb == 14 || nb == 15 || nb == 16 || nb == 17 {
		return 17
	}
	if nb == 18 || nb == 19 {
		return 19
	}
	if nb == 20 || nb == 21 || nb == 22 || nb == 23 {
		return 23
	}
	if nb == 24 || nb == 25 || nb == 26 || nb == 27 || nb == 28 || nb == 29 {
		return 29
	}
	for i := nb; i < 9223372036854775807; i++ {
		pas_bon := 0
		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 || i%11 == 0 || i%13 == 0 || i%17 == 0 || i%19 == 0 || i%23 == 0 || i%29 == 0 {
			pas_bon = 1
		} else {
			for a := nb - 1; a < 29; a-- {
				if a%2 == 0 {
					c = 0
				} else {
					if i%a == 0 {
						pas_bon = 1
					}
				}
			}
		}
		if pas_bon == 0 {
			return i
		}
	}
	return c
}
