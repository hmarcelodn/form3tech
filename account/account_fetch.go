package account

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/config"
)

type AccountFetch struct{}

func (a AccountFetch) Fetch() client.FetchResponse {
	resp, getErr := http.Get(config.AccountURI)

	if getErr != nil {
		log.Fatalln(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatalln(readErr)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalln(resp.StatusCode, string(body))
	}

	var accounts client.FetchResponse
	if jsonErr := json.Unmarshal(body, &accounts); jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	return accounts
}

func (a AccountFetch) FetchByID(id string) client.FetchByIDResponse {
	resp, getErr := http.Get(config.AccountURI + "/" + id)

	if getErr != nil {
		log.Fatalln(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var account client.FetchByIDResponse
	if err := json.Unmarshal(body, &account); err != nil {
		log.Fatalln(err)
	}

	return account
}
