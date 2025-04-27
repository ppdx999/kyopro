package session

import "github.com/ppdx999/kyopro/internal/domain/model"

type SessionAsker interface {
	AskSession() (model.SessionSecret, error)
}

type SessionLoader interface {
	LoadSession() (model.SessionSecret, error)
}

type SessionSaver interface {
	SaveSession(model.SessionSecret) error
}
