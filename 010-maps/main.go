package main

import (
	"fmt"
)

func main() {
	// Map declaration and inicialization
	user := map[string]interface{}{
		"name": map[string]string{
			"firstName": "Felipe",
			"lastName":  "Santos",
			"fullName":  "Felipe da Silva Santos",
		},
		"age": 20,
		"address": map[string]string{
			"publicPlace": "Povoado Oiti",
			"city":        "Limoeiro de Anadia",
			"state":       "Alagoas",
			"country":     "Brasil",
		},
	}

	// Accessing map values
	fmt.Println(user["name"])
	fmt.Println(user["age"])
	fmt.Println(user["address"])

	fmt.Println(user)
	// Deleting a map element
	delete(user, "address")
	fmt.Println(user)

	// Changing a map value
	user["name"] = "Felipe da Silva Santos"

	fmt.Println(user)
}

// OK
