package atcoder

import "github.com/ppdx999/kyopro/internal/domain/model"

type SessionLoader interface {
	LoadSession() (model.SessionSecret, error)
}
