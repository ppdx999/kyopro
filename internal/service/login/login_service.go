package login

/*
LoginServiceはユーザーからSession情報を受け取り、Atcoderにログインして、結果を出力します
*/
type LoginService interface {
	Login() error
}

type LoginServiceImpl struct {
	askSession  AskSession
	loginCheck  LoginCheck
	saveSession SaveSession
	sendMsg     SendMsg
}

func NewLoginServiceImpl(
	askSession AskSession,
	loginCheck LoginCheck,
	saveSession SaveSession,
	sendMsg SendMsg,
) *LoginServiceImpl {
	return &LoginServiceImpl{
		askSession:  askSession,
		loginCheck:  loginCheck,
		saveSession: saveSession,
		sendMsg:     sendMsg,
	}
}

func (l *LoginServiceImpl) Login() error {
	isLogin, err := l.loginCheck.LoginCheck()
	if err != nil {
		return err
	}

	if isLogin {
		l.sendMsg.SendMsg("すでにログインしています")
		return nil
	}

	session, err := l.askSession.AskSession()
	if err != nil {
		return err
	}

	isLogin, err = l.loginCheck.LoginCheck()
	if err != nil || !isLogin {
		l.sendMsg.SendMsg("ログインに失敗しました")
		return err
	}

	err = l.saveSession.SaveSession(session)
	if err != nil {
		l.sendMsg.SendMsg("セッションの保存に失敗しました")
		return err
	}

	l.sendMsg.SendMsg("ログインに成功しました")
	return nil
}
