package main

import (
	"os"

	"github.com/01-edu/z01"
)
func PrintName(s string) string{
	for i := 0; i < len(s); i++ {
		if s[i] != '.' && s[i] != '/'{
			z01.PrintRune(rune(s[i]))
		}
	}
}


func main() {
	argument := os.Args
	text := argument[0]
	PrintName(text)
}
