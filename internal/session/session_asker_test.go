package session_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/session"
)

type MockUserInput struct {
	input string
	err   error
}

func (m *MockUserInput) UserInput() (string, error) {
	return m.input, m.err
}

func TestSessionAsker(t *testing.T) {
	tests := []struct {
		name         string
		userInput    string
		userInputErr error
		want         model.SessionSecret
		wantErr      bool
	}{
		{
			name:      "正常系",
			userInput: "test",
			want:      model.SessionSecret("test"),
		},
		{
			name:         "userinputでエラーが発生",
			userInputErr: errors.New("input error"),
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserInput := &MockUserInput{
				input: tt.userInput,
				err:   tt.userInputErr,
			}
			s := session.NewSessionAskerImpl(mockUserInput)

			got, err := s.AskSession()

			if (err != nil) != tt.wantErr {
				t.Errorf("SessionAsker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SessionAsker() = %v, want %v", got, tt.want)
			}
		})
	}
}
