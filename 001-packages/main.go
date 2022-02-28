package main

import (
	"fmt"               // Built-in package
	"program/auxiliary" // Local package

	"github.com/badoux/checkmail" // External package
)

func main() {
	fmt.Println("Hello, World!")
	// Calls the function of the local package
	auxiliary.Write()

	// Calls the function of the external package
	err := checkmail.ValidateFormat("example@mail.com")
	fmt.Println(err)
}

// OK
