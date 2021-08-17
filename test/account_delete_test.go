package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/client"
)

func TestDeleteAccount(t *testing.T) {
	// Arrange
	accountCreate := account.AccountCreate{}
	accountDelete := account.AccountDelete{}
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

	// Act
	accountCreated, err := accountCreate.Create(account)

	if accountCreated == nil {
		t.Failed()
	}

	if err != nil {
		t.Logf(err.Error())
		t.Failed()
	}

	res, err := accountDelete.Delete(accountCreated.Data.ID)

	if res != false {
		t.Failed()
	}

	// Assert
}

func TestDeleteNonExistingAccount(t *testing.T) {
	accountDelete := account.AccountDelete{}
	nonExistingAccountId, err := uuid.NewRandom()

	if err != nil {
		t.Logf(err.Error())
		t.Failed()
	}

	res, err := accountDelete.Delete(nonExistingAccountId.String())

	if res == false && err == nil {
		t.Failed()
	}
}
