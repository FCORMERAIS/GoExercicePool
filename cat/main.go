package main

import (
    "io/ioutil"
    "os"

    "github.com/01-edu/z01"
)

func PrintStr(s string) {
    for i := 0; i < len(s); i++ {
        z01.PrintRune(rune(s[i]))
    }
}

func main() {
    for i := 1; i < len(os.Args); i++ {
        data, err := ioutil.ReadFile(os.Args[i])
        if err != nil {
            PrintStr("ERROR: ")
            PrintStr(string(err.Error()))
            PrintStr("\n")
            os.Exit(1)
        }
        PrintStr(string(data))
    }
    if len(os.Args) == 1 {
        data, err := ioutil.ReadAll(os.Stdin)
        if err == nil {
            PrintStr(string(data))
        } else {
            PrintStr(err.Error())
        }
    }
}