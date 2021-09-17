package piscine

func BasicAtoi2(s string) int {
	puissanceDix := 1
	intA := 0
	for i := len(s) - 1; i >= 0; i-- {
		intA = intA + int(s[i]%48)*puissanceDix
		puissanceDix = puissanceDix * 10
		if 48 > int(s[i]) || int(s[i]) >= 58 {
			return 0
		}
	}
	return intA
}
