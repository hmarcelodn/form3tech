package test

import (
	"testing"

	"github.com/hmarcelodn/form3tech/account"
)

func TestFetchExistingMultipleAccounts(t *testing.T) {
	accountCreate := account.AccountCreate{}
	accountFetch := account.AccountFetch{}
	accountFixture := AccountFixture{}

	account := accountFixture.Create()
	createResp, err := accountCreate.Create(account)

	if createResp == nil {
		t.Fail()
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	res, err := accountFetch.Fetch()

	if len(res.Data) == 0 {
		t.Logf(`Error: Non rows were fetched.`)
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchExistingSingleAccount(t *testing.T) {
	accountFetch := account.AccountFetch{}
	accountCreate := account.AccountCreate{}
	accountFixture := AccountFixture{}

	account := accountFixture.Create()
	createResp, err := accountCreate.Create(account)

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if _, err := accountFetch.FetchByID(createResp.Data.ID); err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	t.Cleanup(func() {
		Truncate()
	})
}

func TestFetchNonExistingAccountID(t *testing.T) {
	accountFetch := account.AccountFetch{}
	if _, err := accountFetch.FetchByID("fake"); err == nil {
		t.Logf(err.Error())
		t.Fail()
	}
}
