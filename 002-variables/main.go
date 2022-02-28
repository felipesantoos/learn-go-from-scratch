package main

import (
	"fmt"
)

func main() {
	// Variable declaration
	var v1 string = "v1"
	var v2 = "v2"
	v3 := "v3"

	fmt.Println(v1, v2, v3)

	// Variable declaration in a block
	var (
		v4 string = "v4"
		v5        = "v5"
	)

	fmt.Println(v4, v5)

	// Declaring multiple variables on one line
	v6, v7 := "v6", "v7"

	fmt.Println(v6, v7)

	// Exchange of values
	v6, v7 = v7, v6

	fmt.Println(v6, v7)

	// Constant declaration
	const c1 string = "c1"
	const c2 = "c2"

	fmt.Println(c1, c2)

	// Constant declaration in a block
	const (
		c3 string = "c3"
		c4        = "c4"
	)

	fmt.Println(c3, c4)
}

// OK
