package atcoder

import "github.com/ppdx999/kyopro/internal/model"

type SessionLoader interface {
	LoadSession() (model.SessionSecret, error)
}
