package main

import (
	"fmt"
)

type person struct {
	ID   int
	name string
	age  int
}

type studant struct {
	person   // Inheritance
	course   string
	subjects []subjects // Composition
}

type subjects struct {
	ID   int
	name string
}

func main() {
	s := studant{
		// Inheritance
		person: person{
			ID:   1,
			name: "Felipe da Silva Santos",
			age:  20,
		},
		course: "Sistemas de Informação",
		// Composition
		subjects: []subjects{
			{
				ID:   1,
				name: "Linguagem de Programação",
			},
			{
				ID:   2,
				name: "Arquitetura e Organização de Computadores",
			},
		},
	}

	fmt.Println(s.ID)          // Inherited
	fmt.Println(s.name)        // Inherited
	fmt.Println(s.age)         // Inherited
	fmt.Println(s.course)      // His own
	fmt.Println(s.subjects[0]) // Composed by
	fmt.Println(s.subjects[1]) // Composed by
}

// OK
