package main

import "fmt"

// A function that accepts another function (callback) as a parameter
func doMath(a int, b int, callback func(int, int)) {
	fmt.Println("Inside doMath function")
	callback(a, b) // calling the callback function
}

// A function to be passed as a callback
func add(x int, y int) {
	fmt.Println("Addition:", x+y)
}

func main() {
	// Passing the 'add' function as a callback to 'doMath'
	doMath(10, 20, add)

	// Using anonymous function as callback
	doMath(5, 3, func(x, y int) {
		fmt.Println("Multiplication:", x*y)
	})
}
