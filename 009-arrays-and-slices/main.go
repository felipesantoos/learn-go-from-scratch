package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Arrays in a block
	var (
		a1 [3]int = [3]int{1, 2, 3}
		a2        = [3]int{4, 5, 6}
		a3        = [...]int{7, 8, 9}
	)

	// Arrays with gopher
	a4 := [3]int{10, 11, 12}
	a5 := [5]int{0: 13, 2: 15}

	// Arrays assignment
	a1[0], a1[1], a1[2], a5[1] = -1, -2, -3, 14

	// Print multiple variables
	fmt.Println(a1, a2, a3, a4, a5)

	// Slices in a block
	var (
		s1 []int = []int{1, 2, 3}
		s2       = []int{4, 5, 6}
	)

	// Slice with gopher
	s3 := []int{7, 8, 9}

	// Slice elements assignment
	s1[0], s1[1], s1[2] = -1, -2, -3

	// Add a new element in a slice
	s3 = append(s3, 10)

	// Create a slice from an array
	s4 := a5[1:4]

	fmt.Println(s1, s2, s3, s4)

	// Change array element value
	a5[3] = 16

	fmt.Println(s4)

	// Create a slice with the make function with 3 params
	// An internal array is referencied
	// Overflow doubles the array size
	s5 := make([]int, 5, 6)

	// len() and cap() functions
	fmt.Println(s5, len(s5), cap(s5))

	s5 = append(s5, 1)
	s5 = append(s5, 2)

	fmt.Println(s5, len(s5), cap(s5))

	// Create a slice with the make function with 2 params
	// An internal array is referencied
	// Overflow doubles the array size
	s6 := make([]int, 5)

	fmt.Println(s6, len(s6), cap(s6))

	s6 = append(s6, 1)

	fmt.Println(s6, len(s6), cap(s6))

	// Reflect
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
}

// OK
