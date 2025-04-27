package application_service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	application_service "github.com/ppdx999/kyopro/internal/application/service"
	"github.com/ppdx999/kyopro/internal/domain/model"
	session_mock "github.com/ppdx999/kyopro/internal/domain/service/session/mock"
	user_mock "github.com/ppdx999/kyopro/internal/domain/service/user/mock"
)

func Test_loginer_Login(t *testing.T) {
	type mocks struct {
		sessionAsker *session_mock.MockSessionAsker
		loginChecker *user_mock.MockLoginChecker
		sessionSaver *session_mock.MockSessionSaver
		msgSender    *user_mock.MockMsgSender
	}
	tests := []struct {
		name    string
		mocks   func(c *gomock.Controller) *mocks
		wantErr bool
	}{
		{
			name: "正常系",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					sessionAsker: func() *session_mock.MockSessionAsker {
						m := session_mock.NewMockSessionAsker(c)
						m.EXPECT().AskSession().Return(model.SessionSecret("session"), nil)
						return m
					}(),
					loginChecker: func() *user_mock.MockLoginChecker {
						m := user_mock.NewMockLoginChecker(c)
						m.EXPECT().LoginCheck().Return(false, nil)
						m.EXPECT().LoginCheck().Return(true, nil)
						return m
					}(),
					sessionSaver: func() *session_mock.MockSessionSaver {
						m := session_mock.NewMockSessionSaver(c)
						m.EXPECT().SaveSession(model.SessionSecret("session")).Return(nil)
						return m
					}(),
					msgSender: func() *user_mock.MockMsgSender {
						m := user_mock.NewMockMsgSender(c)
						m.EXPECT().SendMsg("ログインに成功しました")
						return m
					}(),
				}
			},
		},
		{
			name: "ログイン済み",
			mocks: func(c *gomock.Controller) *mocks {
				return &mocks{
					loginChecker: func() *user_mock.MockLoginChecker {
						m := user_mock.NewMockLoginChecker(c)
						m.EXPECT().LoginCheck().Return(true, nil)
						return m
					}(),
					msgSender: func() *user_mock.MockMsgSender {
						m := user_mock.NewMockMsgSender(c)
						m.EXPECT().SendMsg("すでにログインしています")
						return m
					}(),
				}
			},
		},
		// TODO: エラーケースのテスト
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mocks := tt.mocks(mockCtrl)

			l := application_service.NewLoginer(
				mocks.sessionAsker,
				mocks.loginChecker,
				mocks.sessionSaver,
				mocks.msgSender,
			)

			// Act & Assert
			if err := l.Login(); (err != nil) != tt.wantErr {
				t.Errorf("loginer.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
