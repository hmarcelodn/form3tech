package test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/hmarcelodn/form3tech/account"
	"github.com/hmarcelodn/form3tech/utils"
)

func TestDeleteWithFailedRequest(t *testing.T) {
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 400}, errors.New("Mock: Failed Forced Error")
	}

	accountId, err := uuid.NewRandom()
	accountDelete := account.AccountDelete{}
	res, err := accountDelete.Delete(accountId.String())

	if res != false {
		t.Fail()
	}

	if err == nil {
		t.Logf(err.Error())
		t.Fail()
	}
}

func TestDeleteWithBadRequestResponse(t *testing.T) {
	bodyErr := "account_delete_test: Failed request response"
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 400, Body: r}, nil
	}

	accountId, err := uuid.NewRandom()
	accountDelete := account.AccountDelete{}
	res, err := accountDelete.Delete(accountId.String())

	if res != false {
		t.Fail()
	}

	if err == nil {
		t.Logf(err.Error())
		t.Fail()
	}
}

func TestDeleteWithInvalidBody(t *testing.T) {
	bodyErr := ""
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 400, Body: r}, nil
	}

	accountId, err := uuid.NewRandom()
	accountDelete := account.AccountDelete{}
	res, err := accountDelete.Delete(accountId.String())

	if res != false {
		t.Fail()
	}

	if err == nil {
		t.Logf(err.Error())
		t.Fail()
	}
}

func TestDeleteWithSuccess(t *testing.T) {
	bodyErr := ""
	r := ioutil.NopCloser(bytes.NewReader([]byte(bodyErr)))
	account.Client = &utils.MockClient{}
	utils.DoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: r}, nil
	}

	accountId, err := uuid.NewRandom()
	accountDelete := account.AccountDelete{}
	res, err := accountDelete.Delete(accountId.String())

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}

	if res != true {
		t.Logf("account_delete_test: response must be true")
		t.Fail()
	}
}
