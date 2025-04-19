package main

import "fmt"

// ✅ Standard Named Function with Single Return Value
// Takes two integers as input, returns their sum.
func addition(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

// ✅ Function with Multiple Return Values
// Takes two integers and returns both their sum and product.
func getNums(a int, b int) (int, int) {
	add := a + b
	mul := a * b
	return add, mul
}

func main() {
	a := 10
	b := 20

	// Calling the function that returns a single value
	sum := addition(a, b)
	fmt.Println("Sum:", sum)

	// Calling the function that returns multiple values
	p, q := getNums(5, 5)
	fmt.Println("Addition:", p)
	fmt.Println("Multiplication:", q)
}
