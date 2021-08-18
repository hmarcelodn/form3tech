package test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
)

func TestDeleteAccount(t *testing.T) {
	Seed()

	account.Client = &http.Client{}
	accountDelete := account.AccountDelete{}

	res, err := accountDelete.Delete("0C879B45-CEEF-4350-946C-D672CDC43FB6")

	if res == false {
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

func TestDeleteNonExistingAccount(t *testing.T) {
	Seed()

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
