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
	var name []string
	name = append(name, account.Name)
	accountAttributes := model.AccountAttributes{
		Country:                &account.Country,
		BankID:                 account.BankID,
		Bic:                    account.Bic,
		BankIDCode:             account.BankIDCode,
		AccountNumber:          account.AccountNumber,
		Iban:                   account.Iban,
		Name:                   name,
		CustomerID:             account.CustomerID,
		ProcessingService:      account.ProcessingService,
		UserDefinedInformation: account.UserDefinedInformation,
		ValidationType:         account.ValidationType,
		ReferenceMask:          account.ReferenceMask,
		AcceptanceQualifier:    account.AcceptanceQualifier,
	}

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

	req, err := http.NewRequest(http.MethodPost, config.AccountURI, payload)
	res, reqErr := Client.Do(req)

	if reqErr != nil {
		return nil, reqErr
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, errors.New(string(body))
	}

	accountCreateResponse := client.AccountCreateResponse{}

	if err := json.Unmarshal(body, &accountCreateResponse); err != nil {
		return nil, err
	}

	return &accountCreateResponse, nil
}
