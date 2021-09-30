package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if os.Args[1] == "quest8.txt" {
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
		file.Close()
	}
	 if os.Args[1] == "quest8.txt" || os.Args[2] == "quest8T.txt" {
			file, err := os.Open("quest8T.txt")
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
		file.Close()
	}
}
