package session_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func TestSessionPath(t *testing.T) {
	type mock struct {
		home *MockHome
	}

	tests := []struct {
		name    string
		mock    func(c *gomock.Controller) *mock
		want    string
		wantErr bool
	}{
		{
			name: "正常系",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					home: func() *MockHome {
						m := NewMockHome(c)
						m.EXPECT().Home().Return("/home/user", nil)
						return m
					}(),
				}
			},
			want: "/home/user/.local/share/kyopro/session.txt",
		},
		{
			name: "homeでエラーが発生",
			mock: func(c *gomock.Controller) *mock {
				return &mock{
					home: func() *MockHome {
						m := NewMockHome(c)
						m.EXPECT().Home().Return("", errors.New("home error"))
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

			s := session.NewSessionPath(mock.home)

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
