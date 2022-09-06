package accounts

import (
	"errors"
	"strconv"
)

// error's variable
var errNoMoney = errors.New("Can't Draw")

type Accounts struct { // for export structor
	// the rule is same with export function
	// Owner   string // start charator is Upper case this variable can export
	// Balance int

	owner   string // if start Lower case, cant export but it will be private value
	balance int
}

// like constructor
// NewAccounts this is create instance of Accounts Object
func NewAccounts(owner string) *Accounts { // method what create instance

	accounts := Accounts{owner: owner, balance: 0}

	return &accounts
}

// how to create struct's methods
// between func and function name, (someName *structType)
// if i didnt write with pointer, the Object is copyed Object, so real Object is not changed for this methods
// Deposit x ammount on your accounts
func (a *Accounts) Deposit(amount int) {
	a.balance += amount
}

// in go, exception handling is not exist(like try~catch ...etc)
// so we should return error type
// error type is satisfy two type of value , first is noraml error line 37, second is type of nil(almost same with null in js )
// Draw x amount on your accounts
func (a *Accounts) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// String Like toString methods in javascript and java
func (a Accounts) String() string {
	return a.owner + " and " + strconv.Itoa(a.balance)
}