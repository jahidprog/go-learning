package main

import "fmt"

// ✅ Global variable
var a = 10

func main() {
	// ✅ Local variable
	age := 30

	if age > 18 {
		// ✅ Variable Shadowing:
		// This 'a' is a new local variable that shadows the global 'a' inside this block.
		a := 47
		fmt.Println("Shadowed a (inside if-block):", a) // Output: 47
	}

	// ✅ Outside the if-block, the global 'a' is accessible again
	fmt.Println("Global a (outside if-block):", a) // Output: 10
}
