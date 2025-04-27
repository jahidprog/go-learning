// Receiver Function

package main

import "fmt"

// User নামে একটি custom struct type, যার মধ্যে Name ও Age নামের দুটি field আছে
type User struct {
	Name string // member variable / field / property
	Age  int
}

// সাধারণ function → struct কে parameter হিসেবে নেয়
func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// Receiver Function → struct এর সাথে যুক্ত function
// 🔹 Receiver Function হচ্ছে এমন একটি ফাংশন যেটা কোনো struct টাইপের instance এর সাথে directly কাজ করে।
// 🔹 struct define না করলে receiver function লেখা সম্ভব না
// 🔹 এই ধরনের function struct এর method হিসেবেও পরিচিত
func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}

// আরেকটি Receiver Function
func (usr User) call(a int) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Extra Value:", a)
}

func main() {
	// user1 struct instance তৈরি
	var user1 User
	user1 = User{
		Name: "Jahid",
		Age:  25,
	}

	// Receiver function call → struct instance dot notation দিয়ে কল করতে হয়
	user1.printDetails()

	// shortcut initialization
	user2 := User{
		Name: "Islam",
		Age:  23,
	}

	user2.call(2223)
}

/*
	===============================
	🧠 MEMORY LAYOUT DURING EXECUTION (RAM)
	===============================

	Go প্রোগ্রাম রান হওয়ার সময় RAM-এ নিচের মতো ভাগে memory allocate হয়:

	1️⃣ Code Segment (Text Segment)
		- এখানে compiled করা instructions থাকে।
		- যেমন: function definition, method definition, ইত্যাদি।
		- read-only segment → execute হয়, modify করা যায় না।
		- যেমন:
			- func main() {}
			- func (u User) printDetails() {}
			- func printUserDetails(u User) {}

	2️⃣ Stack Segment (Function Call Stack)
		- এখানে function call হলে তার local variables এবং execution context রাখা হয়।
		- Last In First Out (LIFO) সিস্টেমে কাজ করে।
		- খুব দ্রুত allocate এবং deallocate হয়।
		- যেমন:
			- main() call হলে → main এর local variable (user1, user2) stack এ store হয়
			- call() function call হলে তার local variable (add) stack এ যায়

	3️⃣ Heap Segment (Dynamic Memory)
		- বড় বা runtime-determined data structures এখানে store হয়
		- যেমন: slices, maps, pointers, interfaces
		- Go runtime automatic garbage collection করে → অব্যবহৃত heap memory cleanup হয়
		- slow but flexible memory allocation

	4️⃣ Data Segment
		📍 Initialized Data:
			- যেখানে initialized global বা package level variable রাখা হয়
			- যেমন: var p = 100

		📍 Uninitialized Data (BSS):
			- যেখানে uninitialized global variable রাখা হয়
			- যেমন: var count int

	5️⃣ Constant Segment
		- যেখানে constant values store হয়
		- যেমন: const a = 10
		- এগুলো read-only → runtime এ modify করা যায় না

	===============================
	📌 Summary:
	===============================
	| Segment         | Stores                          | Example               |
	|------------------|----------------------------------|------------------------|
	| Code/Text        | Function, method definitions     | func main()            |
	| Stack            | Local variables, function calls  | user1 in main()        |
	| Heap             | Runtime allocated objects        | slices, maps           |
	| Data             | Global/package-level vars        | var p = 100            |
	| BSS              | Uninitialized global vars        | var total int          |
	| Constant         | const values                     | const a = 10           |

	⛔️ Note:
	- Stack overflow হয় যদি অনেক recursion বা বড়ো function call হয়
	- Heap memory বেশি use করলে garbage collector workload বাড়ে
*/
