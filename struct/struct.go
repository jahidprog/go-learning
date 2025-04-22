package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	var user1 People
	user1 = People{
		name: "Jahid",
		age:  23,
	}

	user2 := People{
		name: "islam",
		age:  25,
	}

	fmt.Println(user1)
	fmt.Println(user2)
}
