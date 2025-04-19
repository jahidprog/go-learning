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

func main() {
	// Function call with arguments: 45 and 32
	processOperation(45, 32, add)
	processOperation(32, 9, func(x, y int) {
		fmt.Println(x * y)
	})
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
