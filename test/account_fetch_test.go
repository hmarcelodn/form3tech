package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func TestFetchExistingMultipleAccounts(t *testing.T) {
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

	res, err := accountFetch.Fetch()

	if len(res.Data) == 0 {
		t.Logf(`Error: Non rows were fetched.`)
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchExistingSingleAccount(t *testing.T) {
	accountFetch := account.AccountFetch{}
	accountCreate := account.AccountCreate{}

	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
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

	if _, err := accountFetch.FetchByID(createResp.Data.ID); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchNonExistingAccountID(t *testing.T) {
	accountFetch := account.AccountFetch{}
	if _, err := accountFetch.FetchByID("fake"); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}
}
