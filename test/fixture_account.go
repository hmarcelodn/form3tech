package test

import (
	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/client"
)

type FixtureAccount struct{}

func (a FixtureAccount) Create() client.Account {
	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()

	if err != nil {
		panic(err)
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
		BankIDCode:     "GBDSC",
	}

	return account
}

func (a FixtureAccount) CreateInvalid() client.Account {
	organisationId, err := uuid.NewRandom()
	accountId, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	account := client.Account{
		Country:        "",
		Name:           "Hugo Marcelo Del Negro",
		BankID:         "123456",
		Bic:            "NWBKGB22",
		AccountNumber:  "",
		Iban:           "",
		OrganisationID: organisationId.String(),
		AccountId:      accountId.String(),
		BankIDCode:     "GBDSC",
	}

	return account
}
