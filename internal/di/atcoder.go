package di

import (
	"github.com/ppdx999/kyopro/internal/atcoder"
	"github.com/ppdx999/kyopro/internal/requester"
)

func Atcoder() *atcoder.Atcoder {
	var AuthRequester = func() requester.Requester {
		return atcoder.NewAtcoderRequester(SessionLoaderImpl())
	}

	return atcoder.NewAtcoder(AuthRequester())
}
