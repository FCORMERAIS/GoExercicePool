package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintName(s string) {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if s[i] != '.' && [i] != '/' {
			z01.PrintRune(rune(sentence[i]))
		}
	}
}

func main() {
	argument := os.Args
	text := argument[0]
	PrintName(text)
	z01.PrintRune(rune('\n'))
}
