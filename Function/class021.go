// Anonymous function and IIFE function
package main

import "fmt"

func add(a, b int) {
	sum := a + b
	fmt.Println(sum)
}
func main() {
	add(28, 33) // we invoked/call the add function

	// Anonymous Function
	// Immediately invoked function expression
	// IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(3, 4) // IIFE

	age := 12 // expression

	// if-expression
	if age > 10 { // if-block
		fmt.Println("Age is greater than 10")
	}

	mul := func(x int, y int) int { // anonymous function is assigned to a variable
		z := x * y
		return z
	}

	fmt.Println(mul(5, 2))

	// function-expression
	func(p int, q int) { // function block
		c := p - q
		fmt.Println(c)
	}(13, 9)
}

func init() {
	fmt.Println("I'll be Called first")
}
