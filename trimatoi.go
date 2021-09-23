package piscine

func TrimAtoi(s string) int {
	result := 0
	puissanceDix := 1
	negative := 0
	for i := len(s) - 1; i > 0; i-- {
		if int(s[i]) == 45 && i == 0 {
			negative += 1
		}
		if 48 > int(s[i]) || int(s[i]) >= 57 {
			result = result + 0
		} else {
			result = result + int(s[i]%48)*puissanceDix
			puissanceDix = puissanceDix * 10
		}
	}
	if negative == 1 {
		result = -result
		return result
	} else {
		return result
	}
}
