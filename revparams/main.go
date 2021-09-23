package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintName(s string) {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != '.' && sentence[i] != '/' {
			z01.PrintRune(rune(sentence[i]))
		}
	}
}

func main() {
	argument := os.Args
	for i := len(argument) - 1; i < 1; i++ {
		PrintName(argument[i])
		z01.PrintRune(rune('\n'))
	}
}
