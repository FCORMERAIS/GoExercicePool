package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argument := os.Args
	text := argument[0]
	for i := 0; i < len(argument[0]); i++ {
		z01.PrintRune(rune(text[i]))
	}
}
