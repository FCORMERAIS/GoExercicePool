package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open("quest8.txt")
		if err != nil {
			error := err.Error()
			for i := 0; i < len(error); i++ {
				z01.PrintRune(rune(error[i]))
			}
		}
		arr := make([]byte, 14)
		file.Read(arr)
		for i := 0; i < len(string(arr)); i++ {
			z01.PrintRune(rune(string(arr)[i]))
		}
		z01.PrintRune(rune('\n'))
		file.Close()
	}
}
