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
		userInput    string
		userInputErr error
	}
	tests := []struct {
		name    string
		mock    *mock
		want    model.SessionSecret
		wantErr bool
	}{
		{
			name: "正常系",
			mock: &mock{userInput: "test"},
			want: model.SessionSecret("test"),
		},
		{
			name:    "userinputでエラーが発生",
			mock:    &mock{userInputErr: errors.New("input error")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userInput := NewMockUserInput(mockCtrl)
			userInput.EXPECT().UserInput().Return(tt.mock.userInput, tt.mock.userInputErr)
			s := session.NewSessionAskerImpl(userInput)

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
