package login_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service/login"
)

type MockLoginCheck struct {
	calltimes int
	isLogin   []bool
	err       []error
}

func (m *MockLoginCheck) LoginCheck() (bool, error) {
	isLogin := m.isLogin[m.calltimes]
	err := m.err[m.calltimes]
	m.calltimes++
	return isLogin, err
}

type MockAskSession struct {
	session model.SessionSecret
	err     error
}

func (m *MockAskSession) AskSession() (model.SessionSecret, error) {
	return m.session, m.err
}

type MockSaveSession struct {
	savedSessions []model.SessionSecret
	err           error
}

func (m *MockSaveSession) SaveSession(session model.SessionSecret) error {
	m.savedSessions = append(m.savedSessions, session)
	return m.err
}

type MockSendMsg struct {
	sendMsgs []string
}

func (m *MockSendMsg) SendMsg(msg string) {
	m.sendMsgs = append(m.sendMsgs, msg)
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name                string
		firstLoginCheck     bool
		firstLoginCheckErr  error
		askSession          model.SessionSecret
		askSessionErr       error
		secondLoginCheck    bool
		secondLoginCheckErr error
		saveSessionErr      error
		wantLoginResult     bool
		wantErr             bool
	}{
		{
			name:                "正常系",
			firstLoginCheck:     false,
			firstLoginCheckErr:  nil,
			askSession:          model.SessionSecret("session"),
			askSessionErr:       nil,
			secondLoginCheck:    true,
			secondLoginCheckErr: nil,
			saveSessionErr:      nil,
			wantLoginResult:     true,
			wantErr:             false,
		},
		{
			name:               "ログイン済の場合",
			firstLoginCheck:    true,
			firstLoginCheckErr: nil,
			wantLoginResult:    true,
			wantErr:            false,
		},
		{
			name:               "LoginCheckでエラー",
			firstLoginCheck:    false,
			firstLoginCheckErr: errors.New("login check error"),
			wantLoginResult:    false,
			wantErr:            true,
		},
		{
			name:               "AskSessionでエラー",
			firstLoginCheck:    false,
			firstLoginCheckErr: nil,
			askSession:         model.SessionSecret(""),
			askSessionErr:      errors.New("ask session error"),
			wantLoginResult:    false,
			wantErr:            true,
		},
		{
			name:                "2回目のLoginCheckでエラー",
			firstLoginCheck:     false,
			firstLoginCheckErr:  nil,
			askSession:          model.SessionSecret("session"),
			askSessionErr:       nil,
			secondLoginCheck:    false,
			secondLoginCheckErr: errors.New("second login check error"),
			wantLoginResult:     false,
			wantErr:             true,
		},
		{
			name:                "SaveSessionでエラー",
			firstLoginCheck:     false,
			firstLoginCheckErr:  nil,
			askSession:          model.SessionSecret("session"),
			askSessionErr:       nil,
			secondLoginCheck:    true,
			secondLoginCheckErr: nil,
			saveSessionErr:      errors.New("save session error"),
			wantLoginResult:     false,
			wantErr:             true,
		},
	}

	var isLoginSuccess = func(sendmsg MockSendMsg) bool {
		if len(sendmsg.sendMsgs) == 0 {
			return false
		}

		lastMsg := sendmsg.sendMsgs[len(sendmsg.sendMsgs)-1]
		switch lastMsg {
		case "ログインに成功しました":
			return true
		case "すでにログインしています":
			return true
		case "ログインに失敗しました":
			return false
		case "セッションの保存に失敗しました":
			return false
		default:
			panic("unexpected send message: " + lastMsg)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginCheck := &MockLoginCheck{
				isLogin: []bool{tt.firstLoginCheck, tt.secondLoginCheck},
				err:     []error{tt.firstLoginCheckErr, tt.secondLoginCheckErr},
			}
			askSession := &MockAskSession{
				session: tt.askSession,
				err:     tt.askSessionErr,
			}
			saveSession := &MockSaveSession{
				err: tt.saveSessionErr,
			}
			sendMsg := &MockSendMsg{}

			loginService := login.NewLoginServiceImpl(askSession, loginCheck, saveSession, sendMsg)

			err := loginService.Login()

			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			if isLoginSuccess(*sendMsg) != tt.wantLoginResult {
				t.Errorf("Login() login result = %v, wantLoginResult %v", isLoginSuccess(*sendMsg), tt.wantLoginResult)
			}
		})
	}
}
