package main

import (
	"fmt"
)

func main() {
	// Pass by value
	number := 1
	reverseNumber := changeSign(number)
	fmt.Println(number)
	fmt.Println(reverseNumber)

	// Pass by reference
	newNumber := 2
	fmt.Println(newNumber)
	changeSignWithPointer(&newNumber)
	fmt.Println(newNumber)
}

// Pass by value
func changeSign(number int) int {
	return -number
}

// Pass by reference
func changeSignWithPointer(number *int) {
	*number = *number * -1
}

// OK
