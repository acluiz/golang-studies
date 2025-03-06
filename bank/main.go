package main

import (
	"fmt"

	"golang-studies/bank/accounts"
	"golang-studies/bank/customers"
	"golang-studies/bank/handlers"
)


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

func main() {
	holder := customers.Customer{
		Name       : "Ana",
		CPF        : "00000000000",
		Ocuppation : "architect",
	}

	acc := &accounts.CheckingAccount{
		Holder        : holder,
		BranchNumber  : "001",
		AccountNumber : "01234",
		Balance       : 277,
	}

	holder2 := customers.Customer{
		Name       : "Luiz",
		CPF        : "00000000000",
		Ocuppation : "engineer",
	}
	
	acc2 := &accounts.SavingAccount{
		Holder        : holder2,
		BranchNumber  : "001",
		AccountNumber : "67890",
		Balance       : 125.5,
	}

	command := getCommand()

	switch command {
		case 1:
			handlers.HandleDeposit(acc)
		case 2:
			handlers.HandleWithdraw(acc)
		case 3:
			handlers.HandleTransfer(acc, acc2)
		default:
			return
	}
}