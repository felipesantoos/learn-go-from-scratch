package main

import (
	"fmt"
	"time"
)

func main() {
	// For ~ while
	i := 0
	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	// For ~ default
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For ~ each
	list := []int{1, 2, 3, 4, 5}
	for i, elem := range list {
		fmt.Println(i, elem)
		time.Sleep(time.Second / 2)
	}

	// For ~ string
	word := "abc"
	for _, letter := range word {
		fmt.Println(string(letter))
	}

	// For ~ map
	_map := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
	}
	for key, value := range _map {
		fmt.Println(key, value)
	}
}

// OK
