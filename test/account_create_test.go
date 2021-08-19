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
	fixtureAccountCreateResponse := FixtureAccountCreateResponse{}
	fetchByIdResponse := fixtureAccountCreateResponse.Create()
	r := ioutil.NopCloser(bytes.NewReader([]byte(fetchByIdResponse)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	accountCreate := account.AccountCreate{}
	accountFixture := FixtureAccount{}
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
	accountFixture := FixtureAccount{}
	res, err := accountCreate.Create(accountFixture.Create())

	if res != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestCreateAccountWithBadRequestResponse(t *testing.T) {
	bodyErr := "account_create_test: bad request response"
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body:       r,
		}, nil
	}

	accountCreate := account.AccountCreate{}
	accountFixture := FixtureAccount{}
	res, err := accountCreate.Create(accountFixture.Create())

	if res != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}
