package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	fmt.Printf(err.Error())
	arr := make([]byte, 14)
	file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
}
