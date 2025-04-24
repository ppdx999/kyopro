package di

import "github.com/ppdx999/kyopro/internal/domain/service/user"

func LoginChecker() user.LoginChecker {
	return Atcoder()
}

func UserHome() *user.UserHome {
	return &user.UserHome{}
}

func UserInputFromConsole() *user.UserInputFromConsole {
	return &user.UserInputFromConsole{}
}

func MsgSender() user.MsgSender {
	return &user.ConsoleMsgSender{}
}
