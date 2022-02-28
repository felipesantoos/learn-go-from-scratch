package main

import (
	"fmt"
)

func main() {
	sum, sub, mul, div := ssmd(5, 2)
	fmt.Println(sum, sub, mul, div)
}

// Named returns
func ssmd(x, y float64) (sum, sub, mul, div float64) {
	sum = x + y
	sub = x - y
	mul = x * y
	div = x / y

	return
}

// OK
