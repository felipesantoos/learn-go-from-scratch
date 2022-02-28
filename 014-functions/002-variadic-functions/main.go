package main

import (
	"fmt"
)

func main() {
	// Call variadic function
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	// Call variadic function with a fixed param
	fmt.Println(showIntegerList(
		"Algarismos",
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	))
}

// Variadic function
func sum(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}

	return total
}

// Variadic function with a fixed param
func showIntegerList(listName string, integers ...int) string {
	return fmt.Sprintf("%s: %d", listName, integers)
}

// OK
