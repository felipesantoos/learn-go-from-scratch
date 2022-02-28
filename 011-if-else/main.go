package main

import (
	"fmt"
	"math"
)

func main() {
	// Reading the age from terminal
	var age int
	fmt.Scanf("%d", &age)

	// Condicionals
	if age < 13 {
		fmt.Println("Criança.")
	} else if age < 18 {
		fmt.Println("Adolescente.")
	} else if age < 60 {
		fmt.Println("Adulto.")
	} else {
		fmt.Println("Idoso.")
	}

	// if init
	b, c := 3.0, 4.0
	if a := math.Sqrt(math.Pow(b, 2) + math.Pow(c, 2)); a == 5 {
		fmt.Println("É um triângulo pitagórico:", a, b, c)
	} else {
		fmt.Println("Não é um triângulo pitagórico:", a, b, c)
	}
}

// OK
