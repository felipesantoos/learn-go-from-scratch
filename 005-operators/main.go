package main

import (
	"fmt"
)

func main() {
	// Arithmetic
	sum := 1 + 1
	sub := 1 - 1
	mul := 1 * 1
	div := 1 / 1
	mod := 1 % 1

	fmt.Println(sum, sub, mul, div, mod)

	// Assignment
	var x int = 0
	y := 0

	fmt.Println(x, y)

	// Relational
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Logical
	t, f := true, false
	fmt.Println(t && f)
	fmt.Println(t || f)
	fmt.Println(!t, !f)

	// Unary
	i := 0
	i++
	i += 2
	i--
	i -= 2
	i *= 2
	i /= 2
	i %= 2
}

// OK
