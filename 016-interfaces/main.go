package main

import (
	"fmt"
	"math"
)

// Circle and rectangle structs should
// implement this method to be a shape
type shape interface {
	calculateArea() float64
}

// Circle and rectangle can only be
// passed by parameter if they implement
// calculatoArea() float64
func writeArea(s shape) {
	fmt.Printf("%.2f\n", s.calculateArea())
}

// It implements the shape interface
type circle struct {
	radius float64
}

func (c circle) calculateArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// It implements the shape interface
type rectagle struct {
	width  float64
	height float64
}

func (r rectagle) calculateArea() float64 {
	return r.height * r.width
}

func main() {
	c := circle{radius: 10}
	writeArea(c) // Ask for a shape param

	r := rectagle{height: 10, width: 20}
	writeArea(r) // Ask for a shape param
}

// OK
