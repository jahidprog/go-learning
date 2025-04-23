package account

import "fmt"

type Account struct {
	accountNumber string
	holderName    string
	balance       float64
}

func NewAccount(accountNumber, holderName string, initialDeposit float64) *Account {
	return &Account{
		accountNumber: accountNumber,
		holderName:    holderName,
		balance:       initialDeposit,
	}
}

func (a *Account) Deposit(amount float64) {
	if amount < 0 {
		fmt.Println("Invalid deposite amount")
		return
	}
	a.balance += amount
	fmt.Printf("Deposited %.2f successfully\n", amount)
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		fmt.Println("Insufficient balance")
		return
	}

	a.balance -= amount
	fmt.Printf("Withdrawn %0.2f successfully\n", amount)
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) GetDetails() string {
	return fmt.Sprintf("Account Number: %s, Holder: %s, Balance: %f", a.accountNumber, a.holderName, a.balance)
}
