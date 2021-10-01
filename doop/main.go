package main

import (
	"os"
)

func atoi(s string) int {
	puissanceDix := 1
	intA := 0
	negative := 0
	for i := len(s) - 1; i >= 0; i-- {
		if int(s[i]) == 45 && i == 0 {
			negative += 1
		} else if 48 > int(s[i]) || int(s[i]) >= 58 {
			return 0
		} else {
			intA = intA + int(s[i]%48)*puissanceDix
			puissanceDix = puissanceDix * 10
		}
	}
	if negative == 1 {
		intA = -intA
		return intA
	} else {
		return intA
	}
}

func main() {
	if len(os.Args) == 4 {
		if os.Args[2] == "+" {
			ecrire(string(atoi(os.Args[1]) + atoi(os.Args[3])))
			os.Stderr.WriteString("\n")
		}
		if os.Args[2] == "*" {
			ecrire(string(atoi(os.Args[1]) * atoi(os.Args[3])))
			os.Stderr.WriteString("\n")
		}
		if os.Args[2] == "/" {
			ecrire(string(atoi(os.Args[1]) / atoi(os.Args[3])))
			os.Stderr.WriteString("\n")
		}
		if os.Args[2] == "-" {
			ecrire(string(atoi(os.Args[1]) - atoi(os.Args[3])))
			os.Stderr.WriteString("\n")
		}
		if os.Args[2] == "%" {
			ecrire(string(atoi(os.Args[1]) % atoi(os.Args[3])))
			os.Stderr.WriteString("\n")
		}
	}
}

func ecrire(s string) {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		os.Stderr.WriteString(string(sentence[i]))
	}
}
