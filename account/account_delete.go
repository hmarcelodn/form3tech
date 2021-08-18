package account

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/hmarcelodn/form3tech/config"
)

type AccountDelete struct{}

func (a AccountDelete) Delete(id string) (bool, error) {
	req, err := http.NewRequest(http.MethodDelete, config.GetAccountDeleteUri(id), nil)
	resp, err := Client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return false, errors.New(string(body))
	}

	return true, nil
}
