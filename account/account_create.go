package account

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/config"
	"github.com/hmarcelodn/form3tech/model"
)

type AccountCreate struct{}

func (a AccountCreate) Create(account client.Account) {
	accountId, err := uuid.NewRandom()
	organisationId, err := uuid.NewRandom()

	var accountAttributes model.AccountAttributes
	accountBuilder := model.AccountBuilder{}
	accountAttributes = accountBuilder.
		SetCountry(account.Country).
		SetBankID(account.BankID).
		SetBic(account.Bic).
		SetAccountNumber(account.AccountNumber).
		SetIban(account.Iban).
		SetName(account.Name).
		Build()

	accountData := model.AccountData{
		ID:             accountId.String(),
		OrganisationID: organisationId.String(),
		Attributes:     &accountAttributes,
		Type:           config.RecordType,
	}

	accountCreateReq := client.AccountCreateRequest{
		Data: &accountData,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(accountCreateReq)

	resp, err := http.Post(config.AccountURI, "application/json", payload)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln(resp.StatusCode, string(body))
	}
}
