// Clousure
package main

import "fmt"

const a = 10 // const variable declare in global scope

var p = 100 // variable declare in global scope

func outer() func() { // HOR or first-class function
	money := 100
	age := 30

	fmt.Println("Age: ", age)

	show := func() { // function-expression
		money = money + a + p
		fmt.Println(money)
	}

	return show // returning function
}

func call() {
	incr1 := outer() // function expression

	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("===Bank===")
}

/*
	2 Phases:
		1. Compilation Phase (compile time)
		2. Execution Phase (runtime)

	************* Compilation Phase ***************

	// At this stage, Go parses, compiles, and prepares everything for execution

	📦 Code Segment (Text Segment):
		- Function code is stored here: outer(), call(), main(), init(), and the anonymous closure inside outer().
		- const a = 10 is inlined (i.e., directly embedded into instructions).
		- Anonymous closure from `outer()` gets a compiler-generated name and becomes a real function in the binary.

	📦 Data Segment (Global Memory):
		- Global variables like `var p = 100` are allocated here.
		- `const a = 10` is NOT stored here — it's embedded in the code segment during compilation.

	// Compilation output: a binary file (main or ./main)

	Commands:
		go run main.go   // builds + runs (temporary binary)
		go build main.go // builds and saves executable

	************** Execution Phase ******************

	🧠 MEMORY LAYOUT DURING EXECUTION

	1. 🔄 Init Phase
		- Program starts by searching for `init()` function
		- Loads `init()` into the stack
		- Prints: "===Bank==="
		- Pops `init()` from the stack

	2. 🔄 Main Phase
		- Loads `main()` into the stack
		- Calls `call()` → new frame for `call()` is created

	3. 🔄 Call Phase
		- `incr1 := outer()`
			- New frame for `outer()` created
			- Local variables `money=100`, `age=30` stored in `outer` frame
			- Prints: "Age: 30"
			- A closure function is returned
			- Closure captures the *environment* (money, and implicitly can access a, p)

		- `incr1()` called:
			- New stack frame for closure
			- `money = 100 + 10 + 100 = 210` → prints: 210

		- `incr1()` called again:
			- Closure *remembers* the updated `money` (210)
			- `money = 210 + 10 + 100 = 320` → prints: 320

	4. 🔁 New Closure
		- `incr2 := outer()` creates a *new* closure with *new* `money=100`
			- Prints: "Age: 30"

		- `incr2()` called:
			- `money = 100 + 10 + 100 = 210` → prints: 210

		- `incr2()` called again:
			- `money = 210 + 10 + 100 = 320` → prints: 320

	🧱 Stack:
		- Contains stack frames for:
			init() → main() → call() → outer() → closure() calls

	💾 Heap:
		- Closures that **escape** their function (like `show` from `outer`) are moved to the **heap**
		- Each closure keeps its own environment (captured variables like `money`)
		- Garbage Collector (GC) will clean up heap-allocated closures after they’re no longer used

	📦 Final Output:
		===Bank===
		Age: 30
		210
		320
		Age: 30
		210
		320

*/

/*
	*** Binary Executable file *** (demo binary file for understanding)

	main

	0000011101010010101001
	0000011101010010101001
	0000011101010010101001
	...binary representation of all code...
*/
