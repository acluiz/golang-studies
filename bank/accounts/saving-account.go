package accounts

import (
	"errors"
	"fmt"
	"golang-studies/bank/customers"
)

type SavingAccount struct {
	Holder          customers.Customer
	BranchNumber    string
	AccountNumber   string
	TransactionCode string
	Balance         float64
}

func (acc *SavingAccount) deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.Balance += amount

	return nil
}

func (acc *SavingAccount) withdraw(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.Balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.Balance -= amount

	return nil
}

func (acc *SavingAccount) ViewBalance() {
	fmt.Printf("Current balance for %s: %.2f.\n\n", acc.Holder.Name, acc.Balance)
}

func (acc *SavingAccount) Deposit() {
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

func (acc *SavingAccount) Withdraw() {
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