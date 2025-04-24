package session_test

import (
	"errors"
	"testing"

	"github.com/ppdx999/kyopro/internal/domain/service/session"
	"github.com/ppdx999/kyopro/internal/testutil"
)

func TestSessionPath(t *testing.T) {
	tests := []struct {
		name    string
		home    string
		homeErr error
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			home: "/home/user",
			want: "/home/user/.local/share/kyopro/session.txt",
		},
		{
			name:    "homeでエラーが発生",
			homeErr: errors.New("home error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHome := &testutil.MockHome{
				HomeDir: tt.home,
				Err:     tt.homeErr,
			}
			s := session.NewSessionPath(mockHome)

			got, err := s.SessionPath()

			if (err != nil) != tt.wantErr {
				t.Errorf("SessionPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SessionPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
