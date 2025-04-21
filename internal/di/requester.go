package di

import (
	"github.com/ppdx999/kyopro/internal/requester"
)

func InitializeRequester() requester.Requester {
	return requester.NewRequesterImpl()
}
