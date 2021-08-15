package client

import (
	"github.com/hmarcelodn/form3tech/model"
)

type AccountCreateRequest struct {
	Data *model.AccountData `json:"data"`
}
