package main

import (
	"fmt"
)

var number int

// It's possible to have one by file
func init() {
	fmt.Println("Executing init function...")
	number = 10
}

// It's possible to have one package
func main() {
	fmt.Println("Executing main function...")
	fmt.Println(number)
}

// OK
