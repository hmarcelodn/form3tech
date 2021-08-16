package client

import "github.com/hmarcelodn/form3tech/model"

type AccountCreateResponse struct {
	Data  *model.AccountData `json:"data"`
	Links *model.LinkData    `json:"links"`
}
