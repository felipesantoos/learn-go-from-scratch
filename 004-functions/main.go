package main

import (
	"fmt"
)

func main() {
	// It saves the sum return in a variable
	result := sum(2, 1)
	fmt.Println(result)

	// It saves the the func returned by the func in a variable
	f := funcInsideFunc()
	f("function inside function")

	// Two returns
	sum, sub := sumAndSub(2, 1)
	fmt.Println(sum, sub)

	// Ignoring one return
	sum, _ = sumAndSub(4, 2)
	_, sub = sumAndSub(6, 3)

	fmt.Println(sum, sub)
}

// Function declaration
func sum(x int, y int) int {
	return x + y
}

// Function that returns another function
func funcInsideFunc() func(param string) {
	f := func(param string) {
		fmt.Println(param)
	}

	return f
}

// Function with two returns
func sumAndSub(x, y int) (int, int) {
	sum := sum(x, y)
	sub := x - y

	return sum, sub
}
