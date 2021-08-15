package main

import (
	"fmt"

	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func main() {
	fmt.Println("Account:Create")
	var accountCreate account.AccountCreate
	accountCreate.Create(client.Account{
		Country:       "GB",
		Name:          "Pablo Del Negro",
		BankID:        "123456",
		Bic:           "NWBKGB22",
		AccountNumber: "",
		Iban:          "",
	})

	fmt.Println("\nAccount:Fetch")
	var accountFetch account.AccountFetch
	var fetchResp client.FetchResponse = accountFetch.Fetch()
	for i, s := range fetchResp.Data {
		fmt.Println(i, s.ID, *s.Version)
	}

	fmt.Println("\nAccount:Delete")
	var accountDelete account.AccountDelete
	accountDelete.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc42c")
}
