package main

import "fmt"

// processOperation: Higher-order function
// এই ফাংশনটি parameter হিসেবে আরেকটি function (op) গ্রহণ করে এবং সেটিকে call করে
func processOperation(a int, b int, op func(x int, y int)) {
	op(a, b)
}

// add: First-order named function
// এটি দুটি parameter গ্রহণ করে এবং তাদের যোগফল প্রিন্ট করে
func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}

// multiplier: Closure function return করে
// যেই function return হয়, সেটা তার আশেপাশের scope থেকে variable (factor) মনে রাখে
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // factor এর মান মনে রাখে এবং x এর সাথে গুণ করে
	}
}

func main() {
	// এখানে add ফাংশনকে argument হিসেবে পাঠানো হয়েছে processOperation ফাংশনে
	processOperation(45, 32, add)

	// Anonymous function কে সরাসরি পাঠানো হয়েছে (callback হিসেবে)
	processOperation(32, 9, func(x, y int) {
		fmt.Println(x * y)
	})

	// multiplier ফাংশন কল করে twices নামে একটি closure তৈরি করা হয়েছে
	twices := multiplier(2)

	// twices ফাংশনকে কল করলে এটি factor * x রিটার্ন করে (2 * 12 = 24)
	fmt.Println(twices(12)) // Output: 24
}

/*
	1. Parameter vs Argument
		- Parameter: ফাংশন define করার সময় যেই variable গুলো ব্যবহার করি (যেমন: a, b)
		- Argument: ফাংশন কল করার সময় যেই মান গুলো পাঠাই (যেমন: 45, 32)

	2. First-order Function
		- i. Named function: যেমন -> func add(a int, b int)
		- ii. Anonymous function: যেমন -> func(a int, b int) { ... }
		- iii. IIFE (Immediately Invoked Function Expression): যেমন -> func() { ... }()
		- iv. Function expression: function কে variable এর মধ্যে assign করা -> f := func(...) { ... }

	3. Higher-order Function (HOF)
		- যেই ফাংশন অন্য ফাংশনকে parameter হিসেবে নেয় বা return করে
		- যেমন: processOperation()

	4. Callback Function
		- যেই function কে HOF এর মাধ্যমে parameter হিসেবে পাঠানো হয় এবং পরে call করা হয়
		- যেমন: add বা anonymous function -> func(x, y int) { ... }
