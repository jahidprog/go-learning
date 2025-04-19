package main

import "fmt"

// ✅ init function
// This function is automatically called before the main() function.
// You can have multiple init() functions across different files in the same package.
func init() {
	fmt.Println("I'm the init function.")
}

func main() {
	fmt.Println("Main function")
}
