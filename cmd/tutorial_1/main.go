package main

import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// 	var intNum int = 32767 //default 32 or 64 depending on system
// 	intNum = intNum + 1
// 	fmt.Println(intNum)

// 	var floatNum float64 = 12345678.9
// 	fmt.Println(floatNum)

// 	var floatNum32 float32 = 10.1
// 	var intNum32 int32 = 2
// 	var result float32 = floatNum32 + float32(intNum32)
// 	fmt.Println((result))
// }

func main() {
	
	// 1. Variables in Go
	// 1.1 Declaration and Assignment
	// In Go, variables must have their types either explicitly declared or inferred by the compiler. The var keyword is used, and shorthand (:=) can infer the type during assignment.

	// Explicit type declaration
	var x int = 5
	var y string = "Hello, World!"
	
	// Type inference
	// a := 5
	// b := "Hello, Go!"
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)

	// 	1.2 Naming Conventions
	// Go uses camelCase for variable names and PascalCase for exported names (public).

	// Go naming conventions

	userName := "Alice"      // Local variable (camelCase)
	TotalCount := 100        // Exported variable (PascalCase)
	fmt.Println(userName, TotalCount)

	// 	2. Basic Data Types
	// 2.1 Numbers
	// Go has the following numeric types:

	// int (whole numbers)
	// float32 and float64 (decimal numbers)

	    // Integers
		var integerNum int = 42
		fmt.Printf("Integer: %d\n", integerNum)
	
		// Floats
		var floatNum float64 = 3.14
		scientificNotation := 2.5e-3 // 2.5 * 10^-3
		fmt.Printf("Float: %.2f\n", floatNum) // the number after the pointer indicates the number of decimal places
		fmt.Printf("Scientific notation: %.5f\n", scientificNotation)

	// 		Summary of Format Verbs
	// %d: Decimal integer (base 10).
	// %f: Floating-point number with six decimal places (default).
	// %.Nf: Floating-point number with N decimal places.
	// %e: Scientific notation (e.g., 1.23e+06).
	// These format verbs help control how data is displayed, making fmt.Printf a powerful tool for formatted output in Go.

	// 	2.2 Floating-Point Precision
	// Go uses float32 or float64 for floats. There can still be precision limitations:

	result := 1.2 + 2.3
	fmt.Printf("Result: %.20f\n", result) // Display with high precision

	// 	2.3 Arithmetic Operations
	// Go supports basic arithmetic operations:

	a, b := 5, 3

fmt.Printf("Addition: %d\n", a+b)
fmt.Printf("Subtraction: %d\n", a-b)
fmt.Printf("Multiplication: %d\n", a*b)
fmt.Printf("Division: %.2f\n", float64(a)/float64(b)) // Convert to float
fmt.Printf("Integer Division: %d\n", a/b)            // Truncates result
fmt.Printf("Modulo: %d\n", a%b)
fmt.Printf("Exponentiation: %.0f\n", float64(a)*float64(a)*float64(a)) // Go doesn't support **; use math.Pow

}
