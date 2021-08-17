package test

import (
	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/client"
)

type AccountFixture struct{}

func (a AccountFixture) Create() client.Account {
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
	}

	return account
}

func (a AccountFixture) CreateInvalid() client.Account {
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
	}

	return account
}
