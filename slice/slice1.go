package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x[0:8])
}
