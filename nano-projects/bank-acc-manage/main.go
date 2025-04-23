package main

import (
	"fmt"

	"jahidprog.com/nano-projects/account"
)

func main() {
	name := "jahid"

	if !account.ValidName(name) {
		fmt.Println("Invalid Name")
		return
	}

	acc := account.NewAccount("1230-1233", name, 30000)
	fmt.Println(acc.GetDetails())
	acc.Deposit(80000)
	acc.Deposit(2000)
	acc.Deposit(1000)
	acc.Withdraw(7000)
	acc.Withdraw(1000)
	acc.Withdraw(7232)

	fmt.Println(acc.GetBalance())
	fmt.Println(acc.GetDetails())

}
