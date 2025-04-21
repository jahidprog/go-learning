// End of memory
package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}
func main() {
	call()
	fmt.Println(a)
}

func init() {
	fmt.Println("Hello!")
}

/*
	there are two phases in go program
	1. compilation phases
		go have it's own compiler-> after compiling this file with will generate a binary file
		cmd: go build main.go --> will create binay file
		for runing this binary file cmd: ./main

	cmd:
		go build main.go => compile it => main
		./main
		go run main.go --> compile it(go build main.go) --> ./main

	2. execution phases
		execute hoy line by line

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
