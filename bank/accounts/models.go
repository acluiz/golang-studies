package accounts

import "golang-studies/bank/customers"

type Account interface {
	Transfer(amount float64, receiverAcc Account) error
	Withdraw(amount float64) error
	Deposit(amount float64) error
	ViewBalance()
}

type CheckingAccount struct {
	Holder        customers.Customer
	BranchNumber  string
	AccountNumber string
	Balance       float64
}

type SavingAccount struct {
	Holder          customers.Customer
	BranchNumber    string
	AccountNumber   string
	TransactionCode string
	Balance         float64
}
