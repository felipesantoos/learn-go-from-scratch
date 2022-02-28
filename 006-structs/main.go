package main

import (
	"fmt"
)

type user struct {
	ID       int
	name     name
	username string
	emails   []string
	password string
	address  address
}

type name struct {
	firstName string
	lastName  string
	fullName  string
}

type address struct {
	publicPlace string
	number      string
}

func main() {
	// Creating an "object" with the user struct
	u := user{
		ID: 1,
		name: name{
			firstName: "Felipe",
			lastName:  "Santos",
			fullName:  "Felipe da Silva Santos",
		},
		username: "felipesantoos",
		emails:   []string{"felipesantos@gmail.com"},
		password: "********",
		address: address{
			publicPlace: "Povoado Oiti",
			number:      "S/N",
		},
	}

	// Accessing the struct attributes
	fmt.Println(u.ID)
	fmt.Println(u.name.firstName)
	fmt.Println(u.name.lastName)
	fmt.Println(u.name.fullName)
	fmt.Println(u.username)
	fmt.Println(u.emails[0])
	fmt.Println(u.password)
	fmt.Println(u.address.publicPlace)
	fmt.Println(u.address.number)
}

// OK
