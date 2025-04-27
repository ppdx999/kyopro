package application_service

import (
	"github.com/ppdx999/kyopro/internal/domain/service/session"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
)

/*
LoginServiceはユーザーからSession情報を受け取り、Atcoderにログインして、結果を出力します
*/
type loginer struct {
	sessionAsker session.SessionAsker
	loginChecker user.LoginChecker
	sessionSaver session.SessionSaver
	msgSender    user.MsgSender
}

func NewLoginer(
	sessionAsker session.SessionAsker,
	loginChecker user.LoginChecker,
	sessionSaver session.SessionSaver,
	msgSender user.MsgSender,
) *loginer {
	return &loginer{
		sessionAsker: sessionAsker,
		loginChecker: loginChecker,
		sessionSaver: sessionSaver,
		msgSender:    msgSender,
	}
}

func (l *loginer) Login() error {
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
