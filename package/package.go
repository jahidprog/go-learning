package main

import (
	"fmt"
	pkgjahid "go-learning/package/pkgJahid"
)

func main() {
	fmt.Println("Custom Package:")
	sum := pkgjahid.Multiplication(6, 3)
	fmt.Println(sum)

	var a = pkgjahid.Labib
	var b = pkgjahid.Maly

	fmt.Println(a)
	fmt.Println(b)
}
