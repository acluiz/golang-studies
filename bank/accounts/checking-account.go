package accounts

import (
	"errors"
	"fmt"
	"golang-studies/bank/customers"
)

type CheckingAccount struct {
	Holder        customers.Customer
	BranchNumber  string
	AccountNumber string
	Balance       float64
}

func (acc *CheckingAccount) deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.Balance += amount

	return nil
}

func (acc *CheckingAccount) withdraw(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.Balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.Balance -= amount

	return nil
}

func (acc *CheckingAccount) transfer(amount float64, receiverAcc *CheckingAccount) error {
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

func (acc *CheckingAccount) ViewBalance() {
	fmt.Printf("Current balance for %s: %.2f.\n\n", acc.Holder.Name, acc.Balance)
}

func (acc *CheckingAccount) Deposit() {
	fmt.Println("How much you want to deposit?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.deposit(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deposit completed.")
	acc.ViewBalance()
}

func (acc *CheckingAccount) Withdraw() {
	fmt.Println("How much you want to withdraw?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.withdraw(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Withdrawal completed.")
	acc.ViewBalance()
}

func (acc *CheckingAccount) Transfer(receiverAcc *CheckingAccount) {
	fmt.Println("How much you want to transfer?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.transfer(amount, receiverAcc)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transfer completed.")
	
	acc.ViewBalance()
}
