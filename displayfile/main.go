package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open("quest8.txt")
		if err != nil {
			fmt.Printf(err.Error())
		}
		arr := make([]byte, len(os.Args[2]))
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}
