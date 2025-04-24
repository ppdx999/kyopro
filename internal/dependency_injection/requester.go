package di

import (
	"github.com/ppdx999/kyopro/internal/requester"
)

func Requester() requester.Requester {
	return requester.NewRequesterImpl()
}
