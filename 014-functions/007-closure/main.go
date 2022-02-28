package main

import (
	"fmt"
)

func closure() func() {
	t := "Hello, World!"
	f := func() {
		fmt.Println(t)
	}

	return f
}

func main() {
	t := "Olá, Mundo!"
	fmt.Println(t)

	f := closure()
	f() // Hello, World
}

// OK
