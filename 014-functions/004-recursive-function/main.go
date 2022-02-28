package main

import (
	"fmt"
)

func main() {
	// Convert 1 to uint
	i := uint(1)
	for i <= 10 {
		number := fibonacci(i)
		i++
		fmt.Println(number)
	}
}

// Recursive function
func fibonacci(pos uint) uint {
	if pos == 0 || pos == 1 {
		return pos
	} else {
		return fibonacci(pos-1) + fibonacci(pos-2)
	}
}

// OK
