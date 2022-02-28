package main

import (
	"errors" // Errors' default Go package
	"fmt"
)

func main() {
	// Signed integers
	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)

	fmt.Println(i, i8, i16, i32, i64)

	// Unsigned integers
	var (
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
	)

	fmt.Println(ui, ui8, ui16, ui32, ui64)

	// Alias
	var (
		r rune = i32 // rune = int32
		b byte = ui8 // byte = uint8
	)

	fmt.Println(r, b)

	// Decimals
	var (
		float  float32
		double float64
	)

	fmt.Println(float, double)

	// Strings
	var s1 string = "s1"
	var s2 = `2`
	s3 := "s3"

	fmt.Println(s1, s2, s3)

	// Runes
	var c1 rune = 'x'
	var c2 int32 = 'y'
	var c3 int = 'z'
	c1, c2 = c2, c1

	fmt.Println(c1, c2, c3)

	// Zero values
	var (
		az int     // 0
		bz float64 // 0
		cz string  // ""
	)

	fmt.Println(az, bz, cz)

	// Bools
	var (
		b1 bool = true
		b2 bool = false
		b3      = true
		b4 bool // false
	)

	fmt.Println(b1, b2, b3, b4)

	// Error
	var e1 error // nil
	var e2 error = errors.New("e2")

	fmt.Println(e1, e2)
}

// OK
