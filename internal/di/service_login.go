package di

import "github.com/ppdx999/kyopro/internal/service/login"

func InitializeLoginService() login.LoginService {
	var InitializeAskSession = func() login.AskSession {
		return InitializeAskSessionImpl()
	}

	var InitializeLoginCheck = func() login.LoginCheck {
		return InitializeAtcoder()
	}

	var InitializeSaveSession = func() login.SaveSession {
		return InitializeSaveSessionImpl()
	}

	var InitializeSendMsg = func() login.SendMsg {
		return InitializeSendMsgByConsole()
	}

	askSession := InitializeAskSession()
	loginCheck := InitializeLoginCheck()
	saveSession := InitializeSaveSession()
	sendMsg := InitializeSendMsg()
	return login.NewLoginServiceImpl(askSession, loginCheck, saveSession, sendMsg)

}
