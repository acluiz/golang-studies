package main

import (
	"fmt"

	"golang-studies/bank/accounts"
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
	acc := accounts.Account{
		HolderName    : "Ana",
		BranchNumber  : "001",
		AccountNumber : "01234",
		Balance       : 377,
	}
	
	acc2 := accounts.Account{
		HolderName    : "Luiz",
		BranchNumber  : "001",
		AccountNumber : "67890",
		Balance       : 125.5,
	}

	command := getCommand()

	switch command {
		case 1:
			acc.Deposit()
		case 2:
			acc.Withdraw()
		case 3:
			acc.Transfer(&acc2)
		default:
			return
	}
}