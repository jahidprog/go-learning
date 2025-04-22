// Struct

package main

import "fmt"

// A struct is a collection of fields.
// User struct: user এর name এবং age ধারণ করে
type User struct {
	name string // field বা member variable
	age  int    // field বা member variable
}

func main() {
	// Step-by-step initialization
	var user1 User
	user1 = User{ // Instantiation
		name: "Jahid",
		age:  23,
	}

	fmt.Println("Name: ", user1.name)
	fmt.Println("Age: ", user1.age)

	// Short-hand initialization
	user2 := User{
		name: "islam",
		age:  25,
	}

	fmt.Println("Name: ", user2.name)
	fmt.Println("Age: ", user2.age)
}

/*
🧠 2 Phases of Go Program Execution:
	1. Compilation Phase (compile time)
	2. Execution Phase (runtime)

************* Compilation Phase ***************

✅ What happens:
	- `User` struct type compiled and stored in the code segment.
	- `main()` function compiled.
	- Compiler checks for syntax, types, and builds symbol table.
	- If `main()` function is missing → compilation error.

✅ Useful Commands:
	- `go run main.go`   → compiles to a temporary binary & runs it.
	- `go build main.go` → compiles and creates a binary (e.g., `./main`).

************** Execution Phase ******************

🧠 MEMORY LAYOUT DURING EXECUTION:

- Code Segment:
	→ Type definition `User`
	→ Function definition `main()`

- Stack Segment:
	→ A stack frame is created for `main()` function.
	→ Inside that frame:
		• Variable `user1` → stores actual values of `name` and `age`
		• Variable `user2` → also stores actual values directly (value type)

- Heap Segment:
	→ Not used here because:
		• Structs are value types
		• No pointer involved
		• No closure or escape from stack

✅ Execution Flow:
	1. Stack frame for `main()` created
	2. `user1` initialized in two steps (var + assign)
	3. `user2` initialized using shorthand
	4. Prints:
		Name:  Jahid
		Age:  23
		Name:  islam
		Age:  25
	5. Stack frame destroyed after `main()` finishes

✅ Output:
	Name:  Jahid
	Age:  23
	Name:  islam
	Age:  25
*/
