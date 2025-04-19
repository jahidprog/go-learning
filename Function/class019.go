package main

import "fmt"

// 1. Standard Function (Named Function)
// This is a regular, named function called 'add'.
// It takes two integers as parameters, adds them, and prints the result.
func add(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	add(5, 2) // ✅ Calling the named function with arguments 5 and 2
}
