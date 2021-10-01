package main

import "os"

func Atoi(s string) int {
	puissanceDix := 1
	intA := 0
	b := 0
	negative := 0
	c := 0
	if len(s) == 0 {
		return 0
	}
	if s[0] == 45 {
		negative += 1
		c = 1
	} else if s[0] == 43 {
		c = 1
	}
	for i := len(s) - 1; i >= c; i-- {
		if s[i]%48 < 10 && s[i]/48 == 1 {
			intA = intA + int(s[i]%48)*puissanceDix
			puissanceDix = puissanceDix * 10
		} else {
			b = 1
		}
	}
	if b != 1 {
		if negative == 1 {
			intA = -intA
			return intA
		} else {
			return intA
		}
	} else {
		return 0
	}
}

func PrintNbr(n int) {
	nb := n
	if nb < 0 {
		os.Stdout.WriteString("-")
	}
	tabNumber := make([]int, 1)
	var rest int = 0
	for nb > 10 || nb < -10 {
		rest = nb % 10
		tabNumber = append(tabNumber, rest)
		nb = nb / 10
	}
	rest = nb % 10
	tabNumber = append(tabNumber, rest)
	for end := len(tabNumber) - 1; end > 0; end-- {
		if nb > 0 {
			os.Stdout.WriteString(string(48 + tabNumber[end]))
		} else {
			os.Stdout.WriteString(string(48 - tabNumber[end]))
		}
	}
}

func main() {
	if len(os.Args) == 4 {
		t := 0
		y := 0
		for i := 0; i < len(os.Args[1]); i++ {
			if os.Args[1][i] <= byte(9) {
				t++
			}
		}
		for i := 0; i < len(os.Args[3]); i++ {
			if os.Args[1][i] <= byte(9) {
				y++
			}
		}
		if len(os.Args) == 47 && t == len(os.Args[1]) && y == len(os.Args[3]) {
			a := Atoi(os.Args[1])
			b := Atoi(os.Args[3])
			if os.Args[2] == "+" {
				str := a + b
				PrintNbr(str)
			} else if os.Args[2] == "-" {
				str := a - b
				PrintNbr(str)
			} else if os.Args[2] == "/" {
				if os.Args[3] == "0" {
					os.Stdout.WriteString("no division by 0")
				} else {
					str := a / b
					PrintNbr(str)
				}
			} else if os.Args[2] == "*" {
				str := a * b
				PrintNbr(str)
			} else if os.Args[2] == "%" {
				if os.Args[3] == "0" {
					os.Stdout.WriteString("no modulo by 0")
				} else {
					str := a % b
					PrintNbr(str)
				}
			}
		} else {
			os.Stderr.WriteString("")
		}
		os.Stdout.WriteString("\n")
	}
}
