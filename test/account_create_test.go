package test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/utils"
)

func TestCreateAccountWithSuccess(t *testing.T) {
	bodyErr := "{\n    \"data\": {\n        \"attributes\": {\n            \"alternative_names\": null,\n            \"bank_id\": \"400300\",\n            \"bank_id_code\": \"GBDSC\",\n            \"base_currency\": \"GBP\",\n            \"bic\": \"NWBKGB22\",\n            \"country\": \"GB\",\n            \"name\": [\n                \"Marcelo\"\n            ]\n        },\n        \"created_on\": \"2021-08-18T03:10:10.177Z\",\n        \"id\": \"ad27e265-9605-4b4b-a0e5-3003ea9cc42d\",\n        \"modified_on\": \"2021-08-18T03:10:10.177Z\",\n        \"organisation_id\": \"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c\",\n        \"type\": \"accounts\",\n        \"version\": 0\n    },\n    \"links\": {\n        \"self\": \"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc42d\"\n    }\n}"
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	accountCreate := account.AccountCreate{}
	accountFixture := AccountFixture{}
	res, err := accountCreate.Create(accountFixture.Create())

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if res == nil {
		t.Fail()
	}
}

func TestCreateAccountWithFailedRequest(t *testing.T) {
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("account_create_test: Failed forced error")
	}
	accountCreate := account.AccountCreate{}
	accountFixture := AccountFixture{}
	res, err := accountCreate.Create(accountFixture.Create())

	if res != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}
