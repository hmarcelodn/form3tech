package account

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hmarcelodn/form3tech/config"
)

type AccountDelete struct{}

func (a AccountDelete) Delete(uuid string) {
	req, err := http.NewRequest(http.MethodDelete, config.AccountURI+"/"+uuid+"?"+config.RecordVersion, nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
