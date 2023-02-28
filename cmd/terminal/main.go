package main

import (
	"fmt"

	a "github.com/blackflagsoftware/prognos/internal/entities/account"
)

func main() {
	account := a.Account{AccountName: "TestAccount", OwnerName: "BFS", DateFormat: "yyyy-mm-dd"}
	err := a.Create(account)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", account)
}
