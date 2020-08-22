package main

import (
	"fmt"
	"github.com/Cylowatt/go-tutorial/tutil"
)

// Enumerated constant (enums). But int, and not its own type.
// iota (starts at 0) is scoped to the constant block.
//
// _ is write-only variable. You cannot use it, which is a solution
// to an uninitialised int default 0.
const (
	_ = iota
	a
	b
	c
)

// Explicit.
const (
	error = iota // Solution to uninitialised int default 0.
	cat = iota
	dog = iota
	snake = iota
)

const (
	terror = iota + 50 // Offset iota value. Commonly bit-shifting.
	horror
)

// Useful example.
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// Basically, flags. Can OR them to combine them.
const (
	isUser = 1 << iota
	isAdmin
	isStaff
	isManager
)

func main() {
	tutil.PrintV(a)
	tutil.PrintV(b)
	tutil.PrintV(c)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)

	// Staff and manager. Efficient storage.
	var roles byte = isStaff | isManager
	fmt.Printf("Is manager? %v\n", isManager & roles == isManager)
	fmt.Printf("Is admin? %v\n", isAdmin & roles == isAdmin)
}
