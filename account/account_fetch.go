package account

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarcelodn/form3tech/client"
	"github.com/hmarcelodn/form3tech/config"
)

type AccountFetch struct{}

func (a AccountFetch) Fetch() (*client.FetchResponse, error) {
	resp, getErr := http.Get(config.AccountURI)

	if getErr != nil {
		return nil, getErr
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return nil, readErr
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalln(resp.StatusCode, string(body))
	}

	var accounts client.FetchResponse
	if jsonErr := json.Unmarshal(body, &accounts); jsonErr != nil {
		return nil, jsonErr
	}

	return &accounts, nil
}

func (a AccountFetch) FetchByID(id string) (*client.FetchByIDResponse, error) {
	resp, getErr := http.Get(config.AccountURI + "/" + id)

	if getErr != nil {
		return nil, getErr
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var account client.FetchByIDResponse
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	return &account, nil
}
