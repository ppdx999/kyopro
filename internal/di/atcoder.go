package di

import (
	"github.com/ppdx999/kyopro/internal/atcoder"
	"github.com/ppdx999/kyopro/internal/requester"
)

func Atcoder() *atcoder.Atcoder {
	var AuthRequester = func() requester.Requester {
		sessionLoader := SessionLoaderImpl()
		return atcoder.NewAtcoderRequester(sessionLoader)
	}

	requester := AuthRequester()
	return atcoder.NewAtcoder(requester)
}
