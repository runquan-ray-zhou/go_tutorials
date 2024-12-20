package main

import (
	"fmt"
)

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
	// var x int = 5
	// var y string = "Hello, World!"
	
	// Type inference
	// a := 5
	// b := "Hello, Go!"
	
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(a)
	// fmt.Println(b)

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

	// 3. Booleans
	// 3.1 Truthy and Falsy Values
	// Go has no concept of truthy/falsy. Booleans are explicitly true or false. Other types (like 0, nil, empty strings) do not automatically evaluate to false.

	emptySlice := []int{}
	emptyMap := map[string]string{}

	if len(emptySlice) == 0 {
		fmt.Println("Empty slice is falsy") // Explicit check
	}

	if len(emptyMap) == 0 {
		fmt.Println("Empty map is falsy")
	}

	// 	3.2 Equality Operators
	// Go strictly enforces type matching during comparisons:

	fmt.Println(1 == 1)       // true
	fmt.Println(1 == '1')     // Compilation error: type mismatch
	fmt.Println(1 == true)    // Compilation error: invalid operation

	// 	4. Strings
	// 4.1 String Literals
	// Go supports double quotes (") for string literals. Multiline strings use backticks (`).

	singleLine := "Hello, World!"
	multiLine := `This is a
	multi-line string`

	fmt.Println(singleLine)
	fmt.Println(multiLine)

	// 	4.2 String Indexing and Slicing
	// Strings in Go are immutable and indexed by bytes, not runes (for Unicode).

	text := "Golang"

	// Indexing
	fmt.Printf("First character: %c\n", text[0])
	fmt.Printf("Last character: %c\n", text[len(text)-1])

	// Slicing
	fmt.Println(text[0:3]) // "Gol" (up to, but not including, index 3)

	// Reversing a string requires converting it to a rune slice since Go doesn't provide built-in reverse:

	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes)) // Reverse the string

	// 	4.3 String Formatting
	// Go uses the fmt package for string formatting. Printf is similar to f-strings in Python:

	name := "Alice"
	age := 30

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Expressions
	x, y := 10, 20
	fmt.Printf("The sum of %d and %d is %d\n", x, y, x+y)

	// Arrays
	// var arr1 = [3]int{1,2,3}
	// arr2 := [5]int{4,5,6,7,8}
  
	// fmt.Println(arr1)//[1 2 3]
	// fmt.Println(arr2)//[4 5 6 7 8]

	// var arr1 = [...]int{1,2,3}
	// arr2 := [...]int{4,5,6,7,8}
  
	// fmt.Println(arr1)//[1 2 3]
	// fmt.Println(arr2)//[4 5 6 7 8]
	
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)//[Volvo BMW Ford Mazda]

	// prices := [3]int{10,20,30}

	// fmt.Println(prices[0]) //10
	// fmt.Println(prices[2]) //30

	prices := [3]int{10,20,30}

	prices[2] = 50
	fmt.Println(prices)//[10 20 50]

	// arr1 := [5]int{} //not initialized
	// arr2 := [5]int{1,2} //partially initialized
	// arr3 := [5]int{1,2,3,4,5} //fully initialized
  
	// fmt.Println(arr1) //[0 0 0 0 0]
	// fmt.Println(arr2) //[1 2 0 0 0]
	// fmt.Println(arr3) //[1 2 3 4 5]

	arr1 := [5]int{1:10,2:40} //This example initializes only the second and third elements of the array: 

	fmt.Println(arr1) //[0 10 40 0 0]
	// The len() function is used to find the length of an array:
	// arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// arr2 := [...]int{1,2,3,4,5,6}
  
	// fmt.Println(len(arr1)) //4
	// fmt.Println(len(arr2)) //6
}

	// 	String Methods
	// Use Go's strings package for common string manipulations:
	
	// func string() {
	// 	s := "  Hello, World!  "
	
	// 	fmt.Println(strings.ToUpper(s))
	// 	fmt.Println(strings.ToLower(s))
	// 	fmt.Println(strings.TrimSpace(s))
	// 	fmt.Println(strings.Replace(s, "Hello", "Goodbye", 1))
	
	// 	split := strings.Split(s, ",")
	// 	fmt.Println(split)
	// }