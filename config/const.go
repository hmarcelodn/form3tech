package config

import (
	"fmt"
	"os"
)

var (
	baseURL = os.Getenv("ACCOUNT_SERVICE_ADDR")
)

func GetAccountUri() string {
	return fmt.Sprintf("%s/v1/organisation/accounts", baseURL)
}

func GetAccountByIdUri(id string) string {
	return fmt.Sprintf("%s/v1/organisation/accounts/%s", baseURL, id)
}

func GetAccountDeleteUri(id string) string {
	return fmt.Sprintf("%s/v1/organisation/accounts/%s?version=0", baseURL, id)
}

const (
	RecordType = "accounts"
)
