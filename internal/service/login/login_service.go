package login

/*
LoginServiceはユーザーからSession情報を受け取り、Atcoderにログインして、結果を出力します
*/
type LoginService interface {
	Login() error
}

type LoginServiceImpl struct {
	sessionAsker SessionAsker
	loginChecker LoginChecker
	sessionSaver SessionSaver
	msgSender    MsgSender
}

func NewLoginServiceImpl(
	sessionAsker SessionAsker,
	loginChecker LoginChecker,
	sessionSaver SessionSaver,
	msgSender MsgSender,
) *LoginServiceImpl {
	return &LoginServiceImpl{
		sessionAsker: sessionAsker,
		loginChecker: loginChecker,
		sessionSaver: sessionSaver,
		msgSender:    msgSender,
	}
}

func (l *LoginServiceImpl) Login() error {
	isLogin, err := l.loginChecker.LoginCheck()
	if err != nil {
		return err
	}

	if isLogin {
		l.msgSender.SendMsg("すでにログインしています")
		return nil
	}

	session, err := l.sessionAsker.AskSession()
	if err != nil {
		return err
	}

	err = l.sessionSaver.SaveSession(session)
	if err != nil {
		l.msgSender.SendMsg("セッションの保存に失敗しました")
		return err
	}

	isLogin, err = l.loginChecker.LoginCheck()
	if err != nil || !isLogin {
		l.msgSender.SendMsg("ログインに失敗しました")
		return err
	}

	l.msgSender.SendMsg("ログインに成功しました")
	return nil
}
