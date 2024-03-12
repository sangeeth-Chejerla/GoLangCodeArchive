// Module v2
package main

import (
	"fmt"
)

// Account represents a bank account
type Account struct {
    owner    string
    balance  float64
}

// NewAccount creates a new bank account
func NewAccount(owner string, balance float64) *Account {
    return &Account{owner: owner, balance: balance}
}

// Deposit adds funds to the account
func (a *Account) Deposit(amount float64) {
    a.balance += amount
    fmt.Printf("Deposited %.2f into account of %s\n", amount, a.owner)
}

// Withdraw removes funds from the account
func (a *Account) Withdraw(amount float64) {
    if a.balance >= amount {
        a.balance -= amount
        fmt.Printf("Withdrawn %.2f from account of %s\n", amount, a.owner)
    } else {
        fmt.Printf("Insufficient funds for withdrawal from account of %s\n", a.owner)
    }
}

// GetBalance returns the current balance of the account
func (a *Account) GetBalance() float64 {
    return a.balance
}

func main() {
    // Creating a new account using NewAccount function
    account := NewAccount("John Doe", 1000.0)

    // Performing transactions
    account.Deposit(500.0)
    account.Withdraw(200.0)
    account.Withdraw(1500.0) // Insufficient funds

    // Getting the balance
    fmt.Printf("Balance of account %s: %.2f\n", account.owner, account.GetBalance())
}
