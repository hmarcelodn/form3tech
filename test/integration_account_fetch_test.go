package test

import (
	"testing"

	"github.com/hmarcelodn/form3tech/account"
)

func TestFetchExistingMultipleAccounts(t *testing.T) {
	Seed()

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.Fetch()

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if len(res.Data) < 2 {
		t.Logf(`Error: Non rows were fetched.`)
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchExistingSingleAccount(t *testing.T) {
	Seed()

	accountFetch := account.AccountFetch{}

	if _, err := accountFetch.FetchByID("0C879B45-CEEF-4350-946C-D672CDC43FB5"); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchNonExistingAccountID(t *testing.T) {
	Seed()

	accountFetch := account.AccountFetch{}
	if _, err := accountFetch.FetchByID("fake"); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}
