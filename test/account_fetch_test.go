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

func TestFetchWithValidResponse(t *testing.T) {
	bodyErr := "{\n    \"data\": [\n        {\n            \"attributes\": {\n                \"alternative_names\": null,\n                \"bank_id\": \"400300\",\n                \"bank_id_code\": \"GBDSC\",\n                \"bic\": \"NWBKGB22\",\n                \"country\": \"GB\",\n                \"name\": [\n                    \"Pablo Del Negro\"\n                ]\n            },\n            \"created_on\": \"2021-08-18T04:12:29.647Z\",\n            \"id\": \"298a0b5d-0c56-4ec9-bc07-af953b020109\",\n            \"modified_on\": \"2021-08-18T04:12:29.647Z\",\n            \"organisation_id\": \"74579128-2b3e-4cc2-b0fa-3289c6a182d8\",\n            \"type\": \"accounts\",\n            \"version\": 0\n        }\n    ],\n    \"links\": {\n        \"first\": \"/v1/organisation/accounts?page%5Bnumber%5D=first\",\n        \"last\": \"/v1/organisation/accounts?page%5Bnumber%5D=last\",\n        \"self\": \"/v1/organisation/accounts\"\n    }\n}"
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.Fetch()

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if res == nil {
		t.Fail()
	}
}

func TestFetchWithNoData(t *testing.T) {
	body := "{\n    \"data\": null\n}"
	r := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.Fetch()

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if res.Data != nil {
		t.Fail()
	}
}

func TestFetchWithFailedRequest(t *testing.T) {
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("account_fetch_test: failed request")
	}

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.Fetch()

	if res != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestFetchByIdWithValidResponse(t *testing.T) {
	body := "{\n    \"data\": {\n        \"attributes\": {\n            \"alternative_names\": null,\n            \"bank_id\": \"400300\",\n            \"bank_id_code\": \"GBDSC\",\n            \"bic\": \"NWBKGB22\",\n            \"country\": \"GB\",\n            \"name\": [\n                \"Pablo Del Negro\"\n            ]\n        },\n        \"created_on\": \"2021-08-18T04:27:42.889Z\",\n        \"id\": \"2ef27ff1-498a-4277-8235-56d43cf9740a\",\n        \"modified_on\": \"2021-08-18T04:27:42.889Z\",\n        \"organisation_id\": \"431cd5d4-9750-4eb6-bc64-67b582ef0671\",\n        \"type\": \"accounts\",\n        \"version\": 0\n    },\n    \"links\": {\n        \"self\": \"/v1/organisation/accounts/2ef27ff1-498a-4277-8235-56d43cf9740a\"\n    }\n}"
	r := ioutil.NopCloser(bytes.NewReader([]byte(body)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.FetchByID("2ef27ff1-498a-4277-8235-56d43cf9740a")

	if res == nil {
		t.Fail()
	}

	if res.Data.ID != "2ef27ff1-498a-4277-8235-56d43cf9740a" {
		t.Fail()
	}

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
}

func TestFetchByIdWith404Response(t *testing.T) {
	bodyErr := "{\n    \"error_message\": \"record 2ef27ff1-498a-4277-8235-56d43cf9740a does not exist\"\n}"
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Body:       r,
		}, nil
	}

	accountFetch := account.AccountFetch{}
	res, err := accountFetch.FetchByID("2ef27ff1-498a-4277-8235-56d43cf9740a")

	if res != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}

}
