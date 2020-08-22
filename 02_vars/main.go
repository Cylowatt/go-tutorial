package main

import (
	"fmt"
	"github.com/Cylowatt/go-tutorial/tutil"
	"strconv"
	"strings"
)

// Package-level variable.
// You cannot use the := syntax in package-level variables.
var i = 27

// Var block for multiple var declarations.
// Useful for clarity for related variables.
var (
	name string = "Roger"
	age  int    = 34
)

func main() {
	numExample()

	fmt.Println(i, name, age)

	shadowing()

	variableConversion()
}

func numExample() {
	// Explicit type.
	var num int = 30

	// Omit type.
	var num2 = 40

	// Cast type.
	var num3 int16 = 50

	// Constant.
	const num4 = 60

	// Auto-detect type.
	num5 := 70

	// Multi-assign.
	num6, num7 := 80, 90

	fmt.Println(generateNamedNumberString(num, num2, num3, num4, num5, num6, num7))
}

func generateNamedNumberString(values ...interface{}) string {
	str := make([]string, len(values))

	for index, val := range values {
		str[index] = fmt.Sprintf("num%d: %v (%T)", index+1, val, val)
	}

	return strings.Join(str, "\n")
}

func shadowing() {
	fmt.Printf("i = %d\n", i)

	// Cannot declare the same variable in the same scope.
	// However, the variable in the innermost scope takes precedence.
	var i int = 42

	fmt.Printf("i = %d\n", i)
}

func variableConversion() {
	i := 42
	tutil.PrintV(i)


	// Number conversion works as a function.
	// Some data may be lost, e.g. if converting from float to int.
	j := float32(i)
	tutil.PrintV(j)

	// Converts string to a unicode character with value of i.
	k := string(i)
		tutil.PrintV(k)


	// Converts an integer to string.
	// strconv is useful for converting to and from strings.
	l := strconv.Itoa(i)
	tutil.PrintV(l)
}

// pv Prints a value and its type.
func pv(v interface{}) {
	fmt.Printf("%v, %T\n", v, v)
}
