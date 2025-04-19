// Function Expression or Assign function
package main

import "fmt"

func main() {
	// add(6, 3) // you can't call this assign function cause its not created yet. this is undefined
	add := func(a, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(5, 3) // you can all add after creating assign function.
}

func init() {
	fmt.Println("I'll be called first!!")
}
