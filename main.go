package main

import (
	"fmt"

	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/model"
)

func main() {
	fmt.Println("")
	fmt.Println("Account:Create")
	account.Create()

	fmt.Println("")
	fmt.Println("Account:Fetch")
	var fetchResp model.FetchResponse = account.Fetch()
	for i, s := range fetchResp.Data {
		fmt.Println(i, s.Attributes.Name)
	}

	fmt.Println("")
	fmt.Println("Account:Delete")
	account.Delete()
}
