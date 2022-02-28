package main

import (
	"fmt"
)

func main() {
	var (
		v1 int  = 1
		v2 int  = v1  // copy
		v3 *int = &v1 // memory address
	)

	v1++

	// v* = dereferencing
	fmt.Println(v1, v2, *v3) // 2 1 2

}

// OK
