package session_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
	"github.com/ppdx999/kyopro/internal/testutil"
)

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
			mockUserInput := &testutil.MockUserInput{
				Input: tt.userInput,
				Err:   tt.userInputErr,
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
