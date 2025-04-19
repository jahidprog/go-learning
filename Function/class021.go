// Anonymous Functions and IIFE (Immediately Invoked Function Expression)
package main

import "fmt"

func add(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	add(28, 33) // ✅ Calling the regular 'add' function

	// ✅ Anonymous function immediately invoked (IIFE)
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(3, 4) // IIFE - called immediately after being defined

	age := 12 // Variable declaration with value assignment

	// ✅ if-statement (not an expression in Go, technically a statement)
	if age > 10 { // if-block
		fmt.Println("Age is greater than 10")
	}

	// ✅ Anonymous function assigned to a variable (function expression)
	mul := func(x int, y int) int {
		z := x * y
		return z
	}
	fmt.Println(mul(5, 2)) // Calling the function via the variable

	// ✅ IIFE using anonymous function again (subtraction)
	func(p int, q int) {
		c := p - q
		fmt.Println(c)
	}(13, 9)
}

func init() {
	fmt.Println("I'll be called first!") // init() runs before main()
}
