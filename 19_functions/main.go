package main

import "fmt"

func main() {
	fmt.Printf("Variadic Functions\n------------\n")
	variadicFunctionSum(1, 2, 3, 4, 5, 6)

	fmt.Printf("\nReturn Local Variable as a Pointer\n------------\n")
	x := returnLocalVariableAsPointer()
	fmt.Println("Pointer:", x, "value:", *x)

	fmt.Printf("\nNamed Return Values\n------------\n")
	fmt.Println(namedReturnValues())

	fmt.Printf("\nReturn Error\n------------\n")
	handleResultError(0)
	handleResultError(1)

	fmt.Printf("\nAnonymous Functions\n------------\n")
	anonymousFunctions()

	fmt.Printf("\nFunctions as Types\n------------\n")
	functionsAsTypes()
}

func variadicFunctionSum(values ...int) {
	fmt.Println(values)

	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println("Sum:", result)
}

func returnLocalVariableAsPointer() *int {
	x := 5

	// Automatically moved to the heap.
	return &x
}

func namedReturnValues() (result int) {
	// result is now defined.
	result = 15

	// Implicitly returns result.
	return
}

// Multiple return variables are often used for returning errors.
func returnErrorIfZero(value int) (int, error) {
	// Return as soon as possible (defensive programming)
	// to avoid nesting the function in too many scopes.
	if value == 0 {
		return value, fmt.Errorf("provided number is 0")
	}

	return value, nil
}

func handleResultError(value int) {
	if _, err := returnErrorIfZero(value); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("All OK for value", value)
	}
}

func anonymousFunctions() {
	// It is typically better to avoid capturing variables in functions,
	// since they will behave weirdly with async code.
	// Pass them in.
	func() {
		fmt.Println("I am running in an anonymous function!")
	}() // Run right away.

	// The type is func()
	f := func() {
		fmt.Println("I am running in an anonymous function as a variable!")
	}

	f()
}

func functionsAsTypes() {
	// Complicated function type.
	var divide func(float64, float64) (float64, error)

	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}

		return a / b, nil
	}

	fmt.Printf("Function type: %T\n", divide)

	if d, err := divide(5.0, 3.0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", d)
	}
}
