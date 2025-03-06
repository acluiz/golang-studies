package accounts

import (
	"errors"
	"fmt"
)

func (acc *CheckingAccount) Deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid deposit amount")
	}

	acc.Balance += amount

	return nil
}

func (acc *CheckingAccount) Transfer(amount float64, receiverAcc Account) error {
	if (amount <= 0){
		return errors.New("invalid amount")
	}

	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}

	acc.Withdraw(amount)
	receiverAcc.Deposit(amount)

	return nil
}

func (acc *CheckingAccount) ViewBalance() {
	fmt.Printf("Current balance for %s: %.2f.\n", acc.Holder.Name, acc.Balance)
}

func (acc *CheckingAccount) Withdraw(amount float64) error {
	if (amount <= 0){
		return errors.New("invalid withdrawal amount")
	}
	
	if (amount > acc.Balance){
		return errors.New("withdrawal amount is higher than the balance")
	}
	
	acc.Balance -= amount

	return nil
}