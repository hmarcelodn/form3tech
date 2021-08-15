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

func (a AccountCreate) Create(account client.Account) string {
	accountId, err := uuid.NewUUID()
	organisationId, err := uuid.NewUUID()

	var accountAttributes model.AccountAttributes
	accountBuilder := model.GetAccountBuilder(account.Country)
	accountAttributes = accountBuilder.
		SetBankID(account.BankID).               // required. 6 characters
		SetBic(account.Bic).                     // required.
		SetAccountNumber(account.AccountNumber). // optional. 8 characters
		SetIban(account.Iban).                   // optional.
		SetName(account.Name).                   // required.
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

	log.Println(payload)

	resp, err := http.Post(config.AccountURI, "application/json", payload)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
