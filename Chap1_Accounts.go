package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

func main1(){
	// case 1
	// accounts :=  Accounts.Accounts{Owner: "mingyu", Balance:  1000}
	// fmt.Println(accounts)

	// case 2
	accounts := accounts.NewAccounts("mingyu")
	accounts.Deposit(100)

	// err := accounts.WithDraw(200) // how to error handling in go
	// if(err != nil){
	// 	log.Fatalln(err)// the methods of kill programs
	// }
	fmt.Println(accounts)

}