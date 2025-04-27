// Pointer Example in Go
package main

import "fmt"

// Function that accepts a pointer to an array of 3 integers and prints it
func print(nums *[3]int) {
	fmt.Println(nums)
}

// Function that modifies the first element of an array through its pointer
func changeArray(arr *[3]int) {
	arr[0] = 1000
}

type User struct {
	name   string
	age    int
	salary float64
}

func main() {
	// Declaration of an integer variable
	x := 20
	fmt.Println("Initial value of x:", x)

	// Pointer p stores the address of x
	p := &x

	// Dereferencing pointer p and updating the value at that address
	*p = 100
	fmt.Println("After changing value using pointer, x:", x)

	// Printing the address stored in pointer p
	fmt.Println("Address of x:", p)

	// Printing the value at the address p points to (dereferencing)
	fmt.Println("Value at address p:", *p)

	// Declare and initialize an array
	arr := [3]int{1, 2, 3}

	// Pass the address of the array to the print function
	print(&arr)

	// Declare and initialize another array
	nums := [3]int{1, 2, 33}

	// Pass the address of nums to changeArray function which modifies it
	changeArray(&nums)

	// Print the modified array
	fmt.Println("Array after modification:", nums)

	jahid := User{
		name:   "jahid",
		age:    25,
		salary: 0,
	}

	ptr := &jahid

	fmt.Println(ptr.age)
}
