package di

import "github.com/ppdx999/kyopro/internal/service/login"

func InitializeLoginService() login.LoginService {
	var InitializeSessionAsker = func() login.SessionAsker {
		return InitializeSessionAskerImpl()
	}

	var InitializeLoginCheck = func() login.LoginChecker {
		return InitializeAtcoder()
	}

	var InitializeSessionSaver = func() login.SessionSaver {
		return InitializeSessionSaverImpl()
	}

	var InitializeSendMsg = func() login.MsgSender {
		return InitializeConsoleMsgSender()
	}

	askSession := InitializeSessionAsker()
	loginCheck := InitializeLoginCheck()
	saveSession := InitializeSessionSaver()
	sendMsg := InitializeSendMsg()
	return login.NewLoginServiceImpl(askSession, loginCheck, saveSession, sendMsg)

}
