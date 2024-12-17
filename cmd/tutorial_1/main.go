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
	a := 5
	b := "Hello, Go!"
	
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


}
