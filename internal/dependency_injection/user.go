package di

import "github.com/ppdx999/kyopro/internal/domain/service/user"

func UserHome() *user.UserHome {
	return &user.UserHome{}
}

func UserInputFromConsole() *user.UserInputFromConsole {
	return &user.UserInputFromConsole{}
}

func ConsoleMsgSender() *user.ConsoleMsgSender {
	return &user.ConsoleMsgSender{}
}
