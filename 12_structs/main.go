package main

import (
	"fmt"
	"github.com/Cylowatt/go-tutorial/tutil"
	"reflect"
)

// Package-local struct.
type privatePerson struct{}

// Shorter version than the one below.
type PersonShort struct {
	firstName, lastName, city, gender string
	age                               int
}

func main() {
	// Structs are value types.
	fmt.Printf("Create Structs\n------------\n")
	person1, person2 := createStructs()

	fmt.Printf("\nStruct Methods\n------------\n")
	methods(person1, person2)

	fmt.Printf("\nAnonymous Structs\n------------\n")
	anonymousStructs()

	fmt.Printf("\nEmbedding Structs\n------------\n")
	embedding()

	fmt.Printf("\nTagging\n------------\n")
	tagging()
}

func createStructs() (*Person, *Person) {
	// Verbose. Field name syntax. Preferred.
	person1 := Person{
		firstName: "Sally",
		lastName:  "McOsborne",
		city:      "London",
		gender:    "f",
		age:       30,
	}

	// Alternative. Positional syntax. Better to avoid.
	person2 := Person{"Bob", "Oscar", "Hull", "m", 28}

	fmt.Printf("%v\n", person1)
	fmt.Printf("%v\n", person2)

	fmt.Printf("Person 1 lives in %s\n", person1.city)

	return &person1, &person2
}

func methods(p *Person, p2 *Person) {
	fmt.Println(p.greet())
	fmt.Println(p2.greet())

	p.advanceAge()

	fmt.Println(p.greet())

	p.marry(*p2)

	fmt.Println(p.greet())

	p2.marry(*p)

	fmt.Println(p2.greet())
}

func anonymousStructs() {
	somePerson := struct{ name string }{name: "John Peterson"}
	fmt.Println(somePerson)
}

func embedding() {
	bird := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		SpeedKPH: 10,
		CanFly:   false,
	}

	// Note that this is syntactical sugar for direct access.
	fmt.Println(bird.Name, bird.Origin)

	tutil.PrintV(bird)
}

func tagging() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
