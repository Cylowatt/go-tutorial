package main

import "math"

// An interface is a method signature for structs.
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
