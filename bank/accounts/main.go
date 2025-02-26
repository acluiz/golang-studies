package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	HolderName    string
	BranchNumber  string
	AccountNumber string
	Balance       float64
}

func (acc *Account) deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.Balance += amount

	return nil
}

func (acc *Account) withdraw(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.Balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.Balance -= amount

	return nil
}

func (acc *Account) transfer(amount float64, receiverAcc *Account) error {
	if (amount <= 0){
		return errors.New("invalid amount")
	}

	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}

	acc.withdraw(amount)
	receiverAcc.deposit(amount)

	return nil
}

func (acc *Account) Deposit() {
	fmt.Println("How much you want to deposit?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.deposit(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deposit completed.")
	fmt.Printf("New balance for %v: %v.\n", acc.HolderName, acc.Balance)
}

func (acc *Account) Withdraw() {
	fmt.Println("How much you want to withdraw?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.withdraw(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Withdrawal completed.")
	fmt.Printf("New balance for %v: %v.\n", acc.HolderName, acc.Balance)
}

func (acc *Account) Transfer(receiverAcc *Account) {
	fmt.Println("How much you want to transfer?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.transfer(amount, receiverAcc)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transfer completed.")
	
	fmt.Printf("New balance for %v: %v.\n", acc.HolderName, acc.Balance)
	fmt.Printf("New balance for %v: %v.\n", receiverAcc.HolderName, receiverAcc.Balance)
}
