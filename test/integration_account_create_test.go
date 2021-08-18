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
	accountFixture := AccountFixture{}

	res, err := accountCreate.Create(accountFixture.Create())

	if res == nil {
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
	accountFixture := AccountFixture{}
	account := accountFixture.CreateInvalid()

	if _, err := accountCreate.Create(account); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
