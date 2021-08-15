package main

import (
	"fmt"

	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func main() {
	//
	// Create
	//
	fmt.Println("")
	fmt.Println("Account:Create")
	var accountCreate account.AccountCreate
	accountCreate.Create()

	//
	// Fetch
	//
	fmt.Println("")
	fmt.Println("Account:Fetch")
	var accountFetch account.AccountFetch
	var fetchResp client.FetchResponse = accountFetch.Fetch()
	for i, s := range fetchResp.Data {
		fmt.Println(i, s.ID, *s.Version)
	}

	//
	// Delete
	//
	fmt.Println("")
	fmt.Println("Account:Delete")
	var accountDelete account.AccountDelete
	accountDelete.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc42c")
}
