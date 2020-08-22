package main

type Animal struct {
	// `` - tag container. Tags are just strings, and figuring out
	// what to do with them is up to the client or library code.
	//
	// Typically key-values are separated by :
	// values are typically wrapped in ""
	// and tags are typically separated by spaces.
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	// Embed Animal struct into the bird.
	// Like inheritance, but is actually composition.
	// Automatically handled by the compiler.
	// Use interfaces for shared behaviour.
	Animal
	SpeedKPH float32
	CanFly   bool
}
