package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("error : ", err.Error())
	}
	fmt.Println(file.Stat())
	arr := make([]byte, 6)
	file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
}
