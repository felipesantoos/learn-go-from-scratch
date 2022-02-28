package main

import (
	"errors"
	"fmt"
)

// It accepts anything as parameter
func printAnything(value interface{}) {
	fmt.Println(value)
}

type user struct {
	name string
	age  int
}

func main() {
	// Function that accepts any type of variable
	printAnything(1)
	printAnything(1.1)
	printAnything("One")
	printAnything(true)
	printAnything(user{name: "Felipe", age: 20})
	printAnything(errors.New("Bad request"))

	// Interface can leaves to gambiarra
	_map := map[interface{}]interface{}{
		0:                          "Zero",
		"One":                      1,
		true:                       user{name: "Felipe", age: 20},
		user{name: "João", age: 2}: false,
		1.1:                        errors.New("Bad request"),
	}
	fmt.Println(_map)
}

// OK
