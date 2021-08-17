package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func TestCreateAccountWithValidCountry(t *testing.T) {
	var accountCreate account.AccountCreate
	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()
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

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if _, err := accountCreate.Create(account); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestCreateAccountWithEmptyCountry(t *testing.T) {
	var accountCreate account.AccountCreate
	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()
	account := client.Account{
		Country:        "",
		Name:           "Hugo Marcelo Del Negro",
		BankID:         "123456",
		Bic:            "NWBKGB22",
		AccountNumber:  "",
		Iban:           "",
		OrganisationID: organisationId.String(),
		AccountId:      accountId.String(),
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if _, err := accountCreate.Create(account); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
