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

func (a AccountCreate) Create() string {
	// TODO: Depending on the country, other attributes such as bank_id and bic are mandatory
	// TODO: GB, AU, BE, CA, EE, FR, DE, GR, HK, IE, IT, LU, NL, PL, PT, ES, CH, US

	// TODO: Account Creation
	country := "GB"

	accountId, err := uuid.NewUUID()
	organisationId, err := uuid.NewUUID()
	var name []string
	name = append(name, "Marcelo")

	accountAttributes := model.AccountAttributes{
		Country: &country,
		Name:    name,
	}

	accountData := model.AccountData{
		ID:             accountId.String(),
		OrganisationID: organisationId.String(),
		Attributes:     &accountAttributes,
		Type:           config.RecordType,
	}

	accountCreateReq := client.AccountCreateRequest{
		Data: &accountData,
	}
	// END: Account Creation

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

	return string(body)
}
