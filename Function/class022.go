// Function Expression or Assigned Function
package main

import "fmt"

func main() {
	// add(6, 3) // ❌ You can't call 'add' here because it's not defined yet (undefined at this point).

	add := func(a, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(5, 3) // ✅ You can call 'add' after it's assigned as a function.
}

func init() {
	fmt.Println("I'll be called first!!")
}
