package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver method. Used for read operations.
// Value (non-pointer) structs can only use value receiver methods.
// Other methods are not implemented.
func (p Person) greet() string {
	return fmt.Sprintf(
		"Hello, my name is %s %s and I am %s",
		p.firstName,
		p.lastName,
		strconv.Itoa(p.age), // Integer to string.
	)
}

// Pointer receiver method. Used for more optimised reads or write operations.
// Cannot be used by value structs.
func (p *Person) advanceAge() {
	p.age++
}

func (p *Person) marry(spouse Person) {
	if p.gender == "f" {
		p.lastName = spouse.lastName
	}
}
