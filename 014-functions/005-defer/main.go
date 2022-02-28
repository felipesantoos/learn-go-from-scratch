package main

import (
	"fmt"
)

func print(content interface{}) {
	fmt.Println(content)
}

func main() {
	// Defer = adiar
	defer print(1)
	defer print(2)
	print(3)
}

// OK
