package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func main() {
	fmt.Println("Account:Create")
	var accountCreate account.AccountCreate

	accountId, err := uuid.NewRandom()
	organisationId, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	res, err := accountCreate.Create(client.Account{
		AccountId:      accountId.String(),
		OrganisationID: organisationId.String(),
		Country:        "GB",
		Name:           "Pablo Del Negro",
		BankID:         "123456",
		Bic:            "NWBKGB22",
		AccountNumber:  "",
		Iban:           "",
	})

	fmt.Println(res.Data.ID)

	fmt.Println("\nAccount:Fetch")
	var accountFetch account.AccountFetch
	fetchResp, err := accountFetch.Fetch()
	for i, s := range fetchResp.Data {
		fmt.Println(i, s.ID, *s.Version)
	}

	fmt.Println("\nAccount:FetchByID")
	resp, err := accountFetch.FetchByID(res.Data.ID)
	fmt.Println(resp.Data.ID)

	fmt.Println("\nAccount:Delete")
	var accountDelete account.AccountDelete
	for i, s := range fetchResp.Data {
		fmt.Println(i, s.ID, *s.Version)
		accountDelete.Delete(s.ID)
	}
}
