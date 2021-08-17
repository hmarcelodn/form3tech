package test

import (
	"testing"

	"github.com/hmarcelodn/form3tech/account"
)

func TestCreateAccountWithValidCountry(t *testing.T) {
	var accountCreate account.AccountCreate
	accountFixture := AccountFixture{}
	account := accountFixture.Create()

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
