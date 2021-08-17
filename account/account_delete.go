package account

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/hmarcelodn/form3tech/config"
)

type AccountDelete struct{}

func (a AccountDelete) Delete(uuid string) (bool, error) {
	req, err := http.NewRequest(http.MethodDelete, config.AccountURI+"/"+uuid+"?"+config.RecordVersion, nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New(string(body))
	}

	return true, nil
}
