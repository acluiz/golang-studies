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

func (acc *Account) withdraw(withdrawValue float64) error {
	if (withdrawValue <= 0){
		return errors.New("invalid withdrawal value")
	}
	
	if (withdrawValue > acc.balance){
		return errors.New("withdrawal value higher than balance")
	}
	
	acc.balance -= withdrawValue

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

	withdrawValue := 2.0

	err := acc.withdraw(withdrawValue)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(acc)
}