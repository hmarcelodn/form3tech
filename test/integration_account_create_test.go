package test

import (
	"net/http"
	"testing"

	"github.com/hmarcelodn/form3tech/account"
)

func TestCreateAccountWithValidCountry(t *testing.T) {
	Seed()

	account.Client = &http.Client{}
	var accountCreate account.AccountCreate
	accountFixture := FixtureAccount{}
	account := accountFixture.Create()

	res, err := accountCreate.Create(account)

	if res == nil {
		t.Fail()
	}

	if res.Data.ID != account.AccountId ||
		res.Data.OrganisationID != account.OrganisationID ||
		*res.Data.Attributes.Country != account.Country ||
		res.Data.Attributes.Name[0] != account.Name ||
		res.Data.Attributes.BankID != account.BankID ||
		res.Data.Attributes.BankIDCode != account.BankIDCode ||
		res.Data.Attributes.AccountNumber != account.AccountNumber ||
		res.Data.Attributes.Iban != account.Iban {
		t.Fail()
	}

	if res.Data.OrganisationID != account.OrganisationID {
		t.Fail()
	}

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestCreateAccountWithEmptyCountry(t *testing.T) {
	Seed()

	account.Client = &http.Client{}
	var accountCreate account.AccountCreate
	accountFixture := FixtureAccount{}
	account := accountFixture.CreateInvalid()

	if _, err := accountCreate.Create(account); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
