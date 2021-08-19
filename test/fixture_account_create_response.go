package test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/model"
)

type FixtureAccountCreateResponse struct{}

func (f FixtureAccountCreateResponse) Create() string {
	accountData := accountData()

	linkData := model.LinkData{
		First: "1",
		Last:  "1",
		Self:  fmt.Sprintf("/v1/organisation/accounts/%s", accountData.ID),
	}

	fetchResponse := client.AccountCreateResponse{
		Data:  &accountData,
		Links: &linkData,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(fetchResponse)

	return payload.String()
}
