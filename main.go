package main

import "fmt"

func passByValue(a int) {
	a = 999
}
func passByRef(a *int) {
	*a = 222222
}
func main() {
	a := 1111
	passByValue(a)
	fmt.Println(a)
	passByRef(&a)
	fmt.Println(a)

}
