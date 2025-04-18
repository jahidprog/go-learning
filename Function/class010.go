package main

import "fmt"

// Function
func add(num1 int, num2 int) { // num1 = 10, num2 = 20
	var sum int = num1 + num2
	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20

	add(a, b)
	add(8, 9)
}
