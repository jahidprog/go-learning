package main

import "fmt"

// processOperation: Higher-order function, takes a function as argument
func processOperation(a int, b int, op func(x int, y int)) {
	op(a, b)
}

// add: First-order named function, accepts 2 parameters
func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

// multiplier: Returns a function (closure) that multiplies its input by a given factor
// The returned function remembers the factor value when it was created.
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // Multiplies x by the factor
	}
}

func main() {
	// Function call with arguments: 45 and 32
	processOperation(45, 32, add)
	processOperation(32, 9, func(x, y int) {
		fmt.Println(x * y)
	})

	// Creating a multiplier function with a factor of 2
	twices := multiplier(2)

	// Calling the multiplier function with 12, which results in 12 * 2 = 24
	fmt.Println(twices(12)) // Output: 24
}

/*
	1. Parameter vs Argument
		- Parameter: Function define করার সময় variable (a, b)
		- Argument: Function call করার সময় পাঠানো মান (45, 32)

	2. First-order Function
		i. Named function -> func add(a int, b int)
		ii. Anonymous function -> func(a int, b int) { ... }
		iii. IIFE -> func() { ... }()
		iv. Function expression -> f := func(...) { ... }

	3. Higher-order function
		- যেই ফাংশন অন্য ফাংশনকে parameter হিসেবে নেয় বা return করে
		- যেমন: processOperation()
*/
