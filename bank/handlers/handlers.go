package handlers

import (
	"fmt"
	"golang-studies/bank/accounts"
)

func HandleDeposit(acc accounts.Account) {
	fmt.Println("How much you want to deposit?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.Deposit(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deposit completed.")
	acc.ViewBalance()
}

func HandleViewBalance(acc accounts.Account) {
	acc.ViewBalance()
}

func HandleWithdraw(acc accounts.Account) {
	fmt.Println("How much you want to withdraw?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.Withdraw(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Withdrawal completed.")
	acc.ViewBalance()
}

func HandleTransfer(acc accounts.Account, receiverAcc accounts.Account) {
	fmt.Println("How much you want to transfer?")

	var amount float64
	fmt.Scan(&amount)

	err := acc.Transfer(amount, receiverAcc)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transfer completed.")

	acc.ViewBalance()
	receiverAcc.ViewBalance()
}