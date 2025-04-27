package session_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func TestSessionAsker(t *testing.T) {
	type mock struct {
		userInput *MockUserInput
	}
	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) *mock
		want    model.SessionSecret
		wantErr bool
	}{
		{
			name: "正常系",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					userInput: func() *MockUserInput {
						m := NewMockUserInput(c)
						m.EXPECT().UserInput().Return("test", nil)
						return m
					}(),
				}
			},
			want: model.SessionSecret("test"),
		},
		{
			name: "userinputでエラーが発生",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					userInput: func() *MockUserInput {
						m := NewMockUserInput(c)
						m.EXPECT().UserInput().Return("", errors.New("user input error"))
						return m
					}(),
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mock := tt.mock(mockCtrl)
			s := session.NewSessionAsker(mock.userInput)

			// Act
			got, err := s.AskSession()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionAskerImpl.AskSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("SessionAskerImpl.AskSession() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
