package main

import "fmt"

func readInt(prompt string) int {
    var value int
    fmt.Println(prompt)
    fmt.Scan(&value)
    return value
}