package main

import (
	"fmt"
	"log"

	"github.com/kimjayhyun/study-golang/banking"
)

func main() {
	account := banking.NewAccount("jhkim")

	account.Deposit(100)

	err := account.Withdraw(20)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account)
}
