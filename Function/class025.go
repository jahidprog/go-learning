// End of memory
package main

import "fmt"

// Constant value - compile time এ determine হয় এবং পরিবর্তন করা যায় না
const a = 10

// Global variable - package level এ define করা হয়েছে, পুরো ফাইলে access করা যায়
var p = 100

// call(): একটি ফাংশন, যেখানে একটি function expression (anonymous function) define ও call করা হয়েছে
func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6) // 5 + 6 = 11
	add(p, a) // 100 + 10 = 110
}

func main() {
	call()         // call() ফাংশন কল
	fmt.Println(a) // constant 'a' এর মান প্রিন্ট
}

// init(): Go প্রোগ্রামে main এর আগে init() ফাংশন অটো কল হয় (একটি special function)
func init() {
	fmt.Println("Hello!")
}

/*
	Go প্রোগ্রামে দুইটি ধাপ থাকে:

	1. Compilation Phase (কম্পাইলেশন ধাপ)
		- Go এর নিজস্ব কম্পাইলার রয়েছে (go compiler)
		- source file (.go) compile করে একটি binary file তৈরি করে
		- কমান্ড:
			- go build main.go      → এটি compile করে একটি binary তৈরি করবে (main নামে)
			- ./main                → কম্পাইল হওয়া binary ফাইল রান করাবে
			- go run main.go        → এটি ভেতরে ভিতরে আগে go build করে, তারপর ./main রান করে

	2. Execution Phase (এক্সিকিউশন ধাপ)
		- কম্পাইল হওয়ার পর প্রোগ্রাম লাইন বাই লাইন execute হয়
		- প্রথমে init() → তারপর main() execute হয়
*/

/*

┌──────────────────────────────────────── RAM ─────────────────────────────────────────┐
│                                                                                      │
│  ┌────────────┐   ┌────────────────┐   ┌────────────────────────────┐   ┌─────────┐  │
│  │ Code       │   │ Data Segment   │   │         Stack              │   │  Heap   │  │
│  │ Segment    │   │ (Globals)      │   │ (Runtime Call Frames)      │   │ (unused)│  │
│  │            │   │                │   │                            │   │         │  │
│  │ ┌───────┐  │   │ ┌────────────┐ │   │ ┌────────────────────────┐ │   │         │  │
│  │ │ call  │  │   │ │ var p = 100│ │   │ │ init()                 │ │   │         │  │
│  │ ├───────┤  │   │ └────────────┘ │   │ ├────────────────────────┤ │   │         │  │
│  │ │ main  │  │   │ (const a = 10) │   │ │ main()                 │ │   │         │  │
│  │ ├───────┤  │   │  embedded in   │   │ ├────────────────────────┤ │   │         │  │
│  │ │ init  │  │   │  instructions  │   │ │ call()                 │ │   │         │  │
│  │ └───────┘  │   └────────────────┘   │ │ → add(5,6), add(100,10)│ │   │         │  │
│  │ Functions  │                        │ └────────────────────────┘ │   │         │  │
│  │ Loaded     │                        │ (anonymous func on stack)  │   │         │  │
└──────────────────────────────────────────────────────────────────────────────────────┘


*/
