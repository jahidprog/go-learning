package main

import "fmt"

func addition(n1 int, n2 int) int {
	// Function with Return value and types
	// single value return
	sum := n1 + n2

	return sum
}

func getNums(a int, b int) (int, int) {
	// Function with multiple value return
	add := a + b
	mul := a * b

	return add, mul
}

func main() {
	a := 10
	b := 20

	sum := addition(a, b)
	fmt.Println(sum)

	p, q := getNums(5, 5)
	fmt.Println(p)
	fmt.Println(q)

}
