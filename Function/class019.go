package main

import "fmt"

// 1. Standard Function or Named Function
func add(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	add(5, 2)
}
