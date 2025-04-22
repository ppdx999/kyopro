package di

import (
	"github.com/ppdx999/kyopro/internal/atcoder"
	"github.com/ppdx999/kyopro/internal/requester"
)

func InitializeAtcoder() *atcoder.Atcoder {
	var InitializeAuthRequester = func() requester.Requester {
		sessionLoader := InitializeSessionLoaderImpl()
		return atcoder.NewAtcoderRequester(sessionLoader)
	}

	requester := InitializeAuthRequester()
	return atcoder.NewAtcoder(requester)
}
