package di

import (
	"github.com/ppdx999/kyopro/internal/infrastructure/external_service/atcoder"
)

func Atcoder() *atcoder.Atcoder {
	var Requester = func() atcoder.Requester {
		return atcoder.NewRequester(SessionLoader())
	}

	return atcoder.NewAtcoder(Requester())
}
