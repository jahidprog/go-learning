// Array
package main

import "fmt"

func main() {
	// Declaration of an integer array with a fixed size of 3
	var arr1 [3]int
	arr1[0] = 12 // Assign value 12 to index 0
	arr1[1] = 9  // Assign value 9 to index 1
	arr1[2] = 87 // Assign value 87 to index 2
	// arr1[3] = 90 // Error: Can't access index 3 because array indices start from 0 and end at length-1
	fmt.Println(arr1)

	// Declaration and initialization of an integer array with 4 elements
	var arr2 = [4]int{12, 32, 21, 11}
	fmt.Println(arr2)

	// Shorthand syntax for declaring and initializing an array
	arr3 := [2]int{0, 123}
	fmt.Println(arr3)

	// Using [...] allows the compiler to automatically determine the array size based on the number of elements
	arr4 := [...]int{23, 1231, 3, 11, 3, 5, 1, 5, 6}
	fmt.Println(arr4)
}
