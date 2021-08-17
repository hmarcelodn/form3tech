package config

import "os"

var (
	baseURL    = os.Getenv("ACCOUNT_SERVICE_ADDR")
	AccountURI = baseURL + "/v1/organisation/accounts"
)

const (
	RecordVersion = "version=0"
	RecordType    = "accounts"
)
