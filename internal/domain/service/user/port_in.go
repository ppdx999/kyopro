package user

import "github.com/ppdx999/kyopro/internal/domain/model"

type LoginChecker interface {
	LoginCheck() (bool, error)
}

type MsgSender interface {
	SendMsg(string)
}

type UserInput interface {
	UserInput() (string, error)
}

type Home interface {
	Home() (string, error)
}

type Pipeline interface {
	Pipeline() *model.Pipeline
}
