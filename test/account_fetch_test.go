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
	fixtureAccountFetchResponse := FixtureAccountFetchResponse{}
	fixtureFetchResponse := fixtureAccountFetchResponse.Create()
	r := ioutil.NopCloser(bytes.NewReader([]byte(fixtureFetchResponse)))
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
	fixtureAccountFetchResponse := FixtureAccountFetchResponse{}
	body := fixtureAccountFetchResponse.CreateFetchByIdResponseWithNull()
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
	fixtureAccountFetchResponse := FixtureAccountFetchResponse{}
	fixtureFetchByIdResponse := fixtureAccountFetchResponse.CreateFetchByIdResponse()
	r := ioutil.NopCloser(bytes.NewReader([]byte(fixtureFetchByIdResponse)))
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

	if err != nil {
		t.Logf(err.Error())
		t.Fail()
	}
}

func TestFetchByIdWith404Response(t *testing.T) {
	fixtureAccountFetchResponse := FixtureAccountFetchResponse{}
	body := fixtureAccountFetchResponse.CreateFetchByIdResponseWithError("2ef27ff1-498a-4277-8235-56d43cf9740a")
	r := ioutil.NopCloser(bytes.NewReader([]byte(body)))
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
