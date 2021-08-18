package test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
)

func TestDeleteAccount(t *testing.T) {
	account.Client = &http.Client{}
	accountCreate := account.AccountCreate{}
	accountDelete := account.AccountDelete{}
	accountFixture := AccountFixture{}
	account := accountFixture.Create()

	accountCreated, err := accountCreate.Create(account)

	if accountCreated == nil {
		t.Fail()
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	res, err := accountDelete.Delete(accountCreated.Data.ID)

	if res != false {
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestDeleteNonExistingAccount(t *testing.T) {
	accountDelete := account.AccountDelete{}
	nonExistingAccountId, err := uuid.NewRandom()

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	res, err := accountDelete.Delete(nonExistingAccountId.String())

	if res == false && err == nil {
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
