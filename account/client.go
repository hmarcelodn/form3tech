package account

import (
	"net/http"

	"github.com/hmarcelodn/form3tech/utils"
)

var (
	Client utils.HTTPClient
)

func init() {
	Client = &http.Client{}
}
