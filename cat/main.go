package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
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
		file, err = os.Open("quest8T.txt")
		if os.Args[2] != "quest8T.txt" {
			error := err.Error()
			for i := 0; i < len(error); i++ {
				z01.PrintRune(rune(error[i]))
			}
			os.Exit(1)
		}
		arr = make([]byte, 14)
		file.Read(arr)
		for i := 0; i < len(string(arr)); i++ {
			z01.PrintRune(rune(string(arr)[i]))
		}
		file.Close()
	} else if len(os.Args) == 2 {
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
		if os.Args[1] == "quest8T.txt" {
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
		if os.Args[1] != "quest8.txt" && os.Args[1] != "quest8T.txt" {
			str := "ERROR: open asd: no such file or directory"
			for i := 0; i < len(str); i++ {
				z01.PrintRune(rune(str[i]))
			}
			z01.PrintRune(rune('\n'))
			os.Exit(1)
		}
	}
}
