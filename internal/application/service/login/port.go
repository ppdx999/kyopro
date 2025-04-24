package login

import "github.com/ppdx999/kyopro/internal/domain/model"

type SessionAsker interface {
	AskSession() (model.SessionSecret, error)
}

type SessionSaver interface {
	SaveSession(model.SessionSecret) error
}

type LoginChecker interface {
	LoginCheck() (bool, error)
}

type MsgSender interface {
	SendMsg(string)
}
