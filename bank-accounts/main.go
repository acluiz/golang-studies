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
	fmt.Println("3) Transfer")
	fmt.Println("4) Exit")

	var command int
	fmt.Scan(&command)

	return command
}

func (acc *Account) deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.balance += amount

	return nil
}

func (acc *Account) withdraw(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.balance -= amount

	return nil
}

func (acc *Account) transfer(amount float64, receiverAcc *Account) error {
	if (amount <= 0){
		return errors.New("invalid amount")
	}

	if acc.balance < amount {
		return errors.New("insufficient balance")
	}

	acc.withdraw(amount)
	receiverAcc.deposit(amount)

	return nil
}

func execDeposit(acc Account) {
	fmt.Println("How much you want to deposit?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.deposit(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deposit completed.")
	fmt.Printf("New balance for %v: %v.\n", acc.holder_name, acc.balance)
}

func execWithdraw(acc Account) {
	fmt.Println("How much you want to withdraw?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.withdraw(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Withdrawal completed.")
	fmt.Printf("New balance for %v: %v.\n", acc.holder_name, acc.balance)
}

func execTransfer(acc Account, receiverAcc *Account) {
	fmt.Println("How much you want to transfer?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.transfer(amount, receiverAcc)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transfer completed.")
	
	fmt.Printf("New balance for %v: %v.\n", acc.holder_name, acc.balance)
	fmt.Printf("New balance for %v: %v.\n", receiverAcc.holder_name, receiverAcc.balance)
}

func main() {
	acc := Account{
		holder_name: "Ana",
		branchNumber: "001",
		accountNumber: "01234",
		balance: 377,
	}
	
	acc2 := Account{
		holder_name: "Luiz",
		branchNumber: "001",
		accountNumber: "67890",
		balance: 125.5,
	}

	command := getCommand()

	switch command {
		case 1:
			execDeposit(acc)
		case 2:
			execWithdraw(acc)
		case 3:
			execTransfer(acc, &acc2)
		default:
			return
	}
}