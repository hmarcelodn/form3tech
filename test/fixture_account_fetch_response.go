package test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/config"
	"github.com/hmarcelodn/form3tech/model"
)

type FixtureAccountFetchResponse struct{}

func accountData() model.AccountData {
	country := "GB"
	var name []string
	name = append(name, "Test Name")

	accountId, accountIdErr := uuid.NewRandom()
	organisationId, organisationIdErr := uuid.NewRandom()

	if accountIdErr != nil {
		panic(accountIdErr)
	}

	if organisationIdErr != nil {
		panic(organisationIdErr)
	}

	accountAttributes := model.AccountAttributes{
		Country:                &country,
		BankID:                 "400300",
		Bic:                    "NWBKGB22",
		BankIDCode:             "GBDSC",
		AccountNumber:          "123456789",
		Iban:                   "",
		Name:                   name,
		CustomerID:             "",
		ProcessingService:      "",
		UserDefinedInformation: "",
		ValidationType:         "",
		ReferenceMask:          "",
		AcceptanceQualifier:    "",
	}

	accountData := model.AccountData{
		ID:             accountId.String(),
		OrganisationID: organisationId.String(),
		Attributes:     &accountAttributes,
		Type:           config.RecordType,
	}

	return accountData
}

func (f FixtureAccountFetchResponse) Create() string {
	accountData := accountData()
	accountDataArr := make([]*model.AccountData, 1)
	accountDataArr = append(accountDataArr, &accountData)

	linkData := model.LinkData{
		First: "1",
		Last:  "1",
		Self:  fmt.Sprintf("/v1/organisation/accounts/%s", accountData.ID),
	}

	fetchResponse := client.FetchResponse{
		Data:  accountDataArr,
		Links: &linkData,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(fetchResponse)

	return payload.String()
}

func (f FixtureAccountFetchResponse) CreateFetchByIdResponse() string {
	accountData := accountData()
	linkData := model.LinkData{
		First: "1",
		Last:  "1",
		Self:  fmt.Sprintf("/v1/organisation/accounts/%s", accountData.ID),
	}

	fetchResponse := client.FetchByIDResponse{
		Data:  &accountData,
		Links: &linkData,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(fetchResponse)

	return payload.String()
}

func (f FixtureAccountFetchResponse) CreateFetchByIdResponseWithError(id string) string {
	return "{\n    \"error_message\": \"record " + id + " does not exist\"\n}"
}

func (f FixtureAccountFetchResponse) CreateFetchByIdResponseWithNull() string {
	return "{\n    \"data\": null\n}"
}
