package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint8
}

// Changes the user name in your memory address
func (u *user) setName(name string) {
	u.name = name
}

// Changes the user age in your memory address
func (u *user) setAge(age uint8) {
	u.age = age
}

// Get the user name from your memory address
func (u *user) getName() string {
	return u.name
}

// Get the user age from your memory address
func (u *user) getAge() uint8 {
	return u.age
}

func main() {
	// Methods usage
	u := user{}
	u.setName("Felipe da Silva Santos")
	u.setAge(20)
	fmt.Println(u.getName()) // Felipe da Silva Santos
	fmt.Println(u.getAge())  // 20
}

// OK
