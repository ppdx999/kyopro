package di

import (
	"github.com/ppdx999/kyopro/internal/atcoder"
)

func InitializeAtcoder() *atcoder.Atcoder {
	requester := InitializeRequester()
	return atcoder.NewAtcoder(requester)
}
