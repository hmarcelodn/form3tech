package client

import (
	"github.com/hmarcelodn/form3tech/model"
)

type FetchResponse struct {
	Data  []*model.AccountData `json:"data"`
	Links *model.LinkData      `json:"links"`
}
