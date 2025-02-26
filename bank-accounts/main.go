package main

import (
	"errors"
	"fmt"
)

type Account struct {
	holder_name string
	branchNumber  string
	accountNumber string
	balance       float64
}

func getCommand() int {
	fmt.Println("Choose an action:")
	fmt.Println("1) Deposit")
	fmt.Println("2) Withdraw")
	fmt.Println("3) Exit")

	var command int
	fmt.Scan(&command)

	return command
}

func (acc *Account) deposit() error {
	fmt.Println("How much you want to deposit?")

	var amount float64
	fmt.Scan(&amount)

	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.balance += amount

	fmt.Printf("Deposit completed. New balance: %v.\n", acc.balance)

	return nil
}

func (acc *Account) withdraw() error {
	fmt.Println("How much you want to withdraw?")

	var amount float64
	fmt.Scan(&amount)

	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.balance -= amount

	fmt.Printf("Withdrawal completed. New balance: %v.\n", acc.balance)

	return nil
}

func main() {
	acc := &Account{
		holder_name: "Ana",
		branchNumber: "001",
		accountNumber: "01234",
		balance: 125.5,
	}

	command := getCommand()

	if command == 1 {
		acc.deposit()
		return
	}
	
	if command == 2 {
		acc.withdraw()
		return
	}
}