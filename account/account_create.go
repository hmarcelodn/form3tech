package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/config"
	"github.com/hmarcelodn/form3tech/model"
)

type AccountCreate struct{}

func (a AccountCreate) Create(account client.Account) (*client.AccountCreateResponse, error) {
	var accountAttributes model.AccountAttributes
	accountBuilder := model.AccountBuilder{}
	accountAttributes = accountBuilder.
		SetCountry(account.Country).
		SetBankID(account.BankID).
		SetBic(account.Bic).
		SetAccountNumber(account.AccountNumber).
		SetIban(account.Iban).
		SetName(account.Name).
		SetCustomerID(account.CustomerID).
		SetProcessingService(account.ProcessingService).
		SetUserDefinedInformation(account.UserDefinedInformation).
		SetValidationType(account.ValidationType).
		SetReferenceMask(account.ReferenceMask).
		SetAcceptanceQualifier(account.AcceptanceQualifier).
		Build()

	accountData := model.AccountData{
		ID:             account.AccountId,
		OrganisationID: account.OrganisationID,
		Attributes:     &accountAttributes,
		Type:           config.RecordType,
	}

	accountCreateReq := client.AccountCreateRequest{
		Data: &accountData,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(accountCreateReq)

	resp, err := http.Post(config.AccountURI, "application/json", payload)
	accountCreateResponse := client.AccountCreateResponse{}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, errors.New(string(body))
	}

	if err := json.Unmarshal(body, &accountCreateResponse); err != nil {
		return nil, err
	}

	return &accountCreateResponse, nil
}
