package main

import (
	"fmt"
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
			fmt.Println(os.Args[1] + os.Args[3])
		}
		if os.Args[2] == "*" {
			fmt.Println(atoi(os.Args[1]) * atoi(os.Args[3]))
		}
		if os.Args[2] == "/" {
			fmt.Println(atoi(os.Args[1]) / atoi(os.Args[3]))
		}
		if os.Args[2] == "-" {
			fmt.Println(atoi(os.Args[1]) - atoi(os.Args[3]))
		}
		if os.Args[2] == "%" {
			fmt.Println(atoi(os.Args[1]) % atoi(os.Args[3]))
		}
	}
}
