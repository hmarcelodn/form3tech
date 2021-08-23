package test

import (
	"strings"
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

	res, err := accountFetch.FetchByID("0C879B45-CEEF-4350-946C-D672CDC43FB5")

	if err != nil {
		t.Fail()
	}

	if res == nil {
		t.Fail()
	}

	if res.Data.ID != strings.ToLower("0C879B45-CEEF-4350-946C-D672CDC43FB5") ||
		res.Data.OrganisationID != strings.ToLower("D612C78B-7C54-4985-919F-0E393F034E0D") ||
		*res.Data.Attributes.Country != "GB" ||
		res.Data.Attributes.Name[0] != "Marcelo Del Negro" ||
		res.Data.Attributes.BankID != "400300" ||
		res.Data.Attributes.BankIDCode != "GBDSC" ||
		*res.Data.Version != 0 {
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
