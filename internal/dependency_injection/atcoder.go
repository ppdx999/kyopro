package di

import (
	"github.com/ppdx999/kyopro/internal/infrastructure/external_service/atcoder"
	"github.com/ppdx999/kyopro/internal/requester"
)

func Atcoder() *atcoder.Atcoder {
	var AuthRequester = func() requester.Requester {
		return atcoder.NewAtcoderRequester(SessionLoader())
	}

	return atcoder.NewAtcoder(AuthRequester())
}
