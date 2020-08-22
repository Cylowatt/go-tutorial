package main

import (
	"fmt"
	"github.com/Cylowatt/go-tutorial/tutil"
)

func main() {
	/*
	 * Main primitive types.
	 *
	 * string - encodes any UTF-8 character.
	 *
	 * bool
	 *
	 * int (size unspecified, but guaranteed to be at least 32 bits)
	 * int8 int16 int32 int64
	 *
	 * uint (size unspecified, but guaranteed to be at least 32 bits)
	 * uint8 uint16 uint32 uint64 uintptr
	 *
	 * byte - alias for uint8
	 * rune - alias for int32
	 * float32 float64
	 * complex64 complex128
	 */

	n := true
	tutil.PrintV(n)

	binaryOperators()

	floats()

	complexNumbers()

	strings()

	runes()
}

func binaryOperators() {
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 1010 AND 0011 = 0010
	fmt.Println(a | b)  // 1010 OR 0011 = 1011
	fmt.Println(a ^ b)  // 1010 XOR 0011 = 1001
	fmt.Println(a &^ b) // 1010 AND NOT 0011 = 0100 (bitwise clear; 1 if both 0; opposite of OR)
}

func floats() {
	n := 3.14

	// Exponential notation.
	n = 13.7e72
	n = 2.1e14
	tutil.PrintV(n)
}

func complexNumbers() {
	n := 2i
	n = 17.2 + 2.534i
	tutil.PrintV(n)

	// Get real and imaginary parts.
	tutil.PrintV(real(n))
	tutil.PrintV(imag(n))

	// Create a complex function.
	n = complex(5.3, 12.43)
	tutil.PrintV(n)
}

func strings() {
	// Can index into the string.
	// Each entry is a byte, but a single character may be more than one byte.
	s := "this is a string"

	// Convert a string to a slice of bytes.
	byteSlice := []byte(s)

	tutil.PrintV(byteSlice)
}

// A rune (int32 alias) represents a UTF-32 character unlike string's UTF-8.
// Any valid UTF-8 character is also a valid UTF-32 character.
//
// You can use built-in functions like Reader.ReadRune to work with runes.
func runes() {
	r := 'a'
	tutil.PrintV(r)
}
