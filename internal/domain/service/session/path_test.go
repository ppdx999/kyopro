package session_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func TestSessionPath(t *testing.T) {
	type mock struct {
		home    string
		homeErr error
	}

	tests := []struct {
		name    string
		mock    *mock
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			mock: &mock{
				home: "/home/user",
			},
			want: "/home/user/.local/share/kyopro/session.txt",
		},
		{
			name:    "homeでエラーが発生",
			mock:    &mock{homeErr: errors.New("home error")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockHome := NewMockHome(mockCtrl)
			mockHome.EXPECT().Home().Return(tt.mock.home, tt.mock.homeErr)

			s := session.NewSessionPath(mockHome)

			// Act
			got, err := s.SessionPath()

			// Assert
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
