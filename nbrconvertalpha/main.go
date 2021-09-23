package main

import (
	"os"

	"github.com/01-edu/z01"
)

func maj(s string) int {
	if s == "--upper" {
		return 1
	} else {
		return 0
	}
}

func Atoi(s string) int {
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
	argument := os.Args
	if len(argument) > 1 {
		if maj(argument[1]) == 1 {
			for i := 2; i < len(argument); i++ {
				df := Atoi(argument[i])
				if 64+df <= 90 && 64+df >= 65 {
					z01.PrintRune(rune(64 + df))
				} else {
					z01.PrintRune(rune(' '))
				}
			}
		} else {
			for k := 1; k < len(argument); k++ {
				df := Atoi(argument[k])
				if df+96 >= 97 && 96+df <= 122 {
					z01.PrintRune(rune(96 + df))
				} else {
					z01.PrintRune(rune(' '))
				}
			}
		}
		z01.PrintRune('\n')
	}
}
