package piscine

func TrimAtoi(s string) int {
	result := 0
	puissanceDix := 1
	negative := 0
	if len(s) == 0 {
		return 0
	}
	for k := 0; negative == 0 && len(s) > k; k++ {
		if 48 <= int(s[k]) && int(s[k]) <= 57 {
			negative = -9999
		} else if rune(s[k]) == 45 {
			negative += 1
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if 48 >= int(s[i]) || int(s[i]) >= 57 {
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
