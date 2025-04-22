package di

import "github.com/ppdx999/kyopro/internal/service/login"

func LoginService() login.LoginService {
	var SessionAsker = func() login.SessionAsker {
		return SessionAskerImpl()
	}

	var LoginCheck = func() login.LoginChecker {
		return Atcoder()
	}

	var SessionSaver = func() login.SessionSaver {
		return SessionSaverImpl()
	}

	var SendMsg = func() login.MsgSender {
		return ConsoleMsgSender()
	}

	askSession := SessionAsker()
	loginCheck := LoginCheck()
	saveSession := SessionSaver()
	sendMsg := SendMsg()
	return login.NewLoginServiceImpl(askSession, loginCheck, saveSession, sendMsg)

}
