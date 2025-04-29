package main

import "fmt"

func main() {
	// arr := [6]string{"This", "is", "a", "GO", "Interview", "Questions"}
	// fmt.Println(arr)
	// s := arr[1:4]
	// fmt.Println(s)
	// s1 := s[1:2]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// s2 := make([]int, 3)
	// s2[0] = 12

	// fmt.Println(s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	// s3 := make([]int, 3, 5)
	// s3[0] = 23
	// s3[2] = 19
	// s3[3] = 20
	// fmt.Println(s3)
	// fmt.Println(len(s3))
	// fmt.Println(cap(s3))

	var x []int
	x = append(x, 1) // len = 1, cap = 1
	fmt.Println("Len", len(x), "Cap", cap(x))
	x = append(x, 2) // len = 2, cap = 2
	fmt.Println("Len", len(x), "Cap", cap(x))
	x = append(x, 3) // len = 3, cap = 4
	fmt.Println("Len", len(x), "Cap", cap(x))

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x, "Len", len(x), "Cap", cap(x)) // [10, 2, 3, 5] len = 4, cap = 4
	fmt.Println(y, "Len", len(y), "Cap", cap(y)) // [10, 2, 3, 5] len = 4, cap = 4
}
