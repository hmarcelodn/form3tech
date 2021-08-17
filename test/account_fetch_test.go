package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func TestFetchExistingAccount(t *testing.T) {
	accountCreate := account.AccountCreate{}
	accountFetch := account.AccountFetch{}
	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()

	if err != nil {
		t.Logf(err.Error())
		t.Failed()
	}

	account := client.Account{
		Country:        "GB",
		Name:           "Hugo Marcelo Del Negro",
		BankID:         "123456",
		Bic:            "NWBKGB22",
		AccountNumber:  "",
		Iban:           "",
		OrganisationID: organisationId.String(),
		AccountId:      accountId.String(),
	}

	createResp, err := accountCreate.Create(account)

	if createResp == nil {
		t.Fail()
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	fetchResp := accountFetch.Fetch()

	if len(fetchResp.Data) == 0 {
		t.Logf(`Error: Non rows were fetched.`)
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
