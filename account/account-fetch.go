package account

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarcelodn/form3tech/model"
)

func Fetch() model.FetchResponse {
	url := "http://localhost:8080/v1/organisation/accounts"
	resp, getErr := http.Get(url)

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

	var accounts model.FetchResponse
	jsonErr := json.Unmarshal(body, &accounts)

	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	return accounts
}
