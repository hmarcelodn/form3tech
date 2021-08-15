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

	var accounts client.FetchResponse
	jsonErr := json.Unmarshal(body, &accounts)

	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	return accounts
}
