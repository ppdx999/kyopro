package login

import "github.com/ppdx999/kyopro/internal/model"

type AskSession interface {
	AskSession() (model.SessionSecret, error)
}

type SaveSession interface {
	SaveSession(model.SessionSecret) error
}

type LoginCheck interface {
	LoginCheck() (bool, error)
}

type SendMsg interface {
	SendMsg(string)
}
