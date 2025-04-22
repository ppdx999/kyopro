package di

import "github.com/ppdx999/kyopro/internal/user"

func InitializeUserHome() *user.UserHome {
	return &user.UserHome{}
}

func InitializeUserInputFromConsole() *user.UserInputFromConsole {
	return &user.UserInputFromConsole{}
}

func InitializeConsoleMsgSender() *user.ConsoleMsgSender {
	return &user.ConsoleMsgSender{}
}
