package helper_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service/helper"
)

type MockGetWd struct {
	path string
	err  error
}

func (m MockGetWd) GetWd() (string, error) {
	return m.path, m.err
}

func GetWorkspaceTest(t *testing.T) {
	tests := []struct {
		name     string
		getwd    string
		getwdErr error
		wantErr  bool
		want     *model.Workspace
	}{
		{
			name:     "success",
			getwd:    "/user/current/workdir",
			getwdErr: nil,
			wantErr:  false,
			want: &model.Workspace{
				Path: "/user/current/workdir",
			},
		},
		{
			name:     "failed getwd",
			getwd:    "/user/current/workdir",
			getwdErr: errors.New("failed getwd"),
			wantErr:  true,
			want:     &model.Workspace{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGetwd := MockGetWd{
				path: tt.getwd,
				err:  tt.getwdErr,
			}

			service := helper.GetWorkspaceImpl{
				GetWd: mockGetwd,
			}

			got, err := service.GetWorkspace()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkspace() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
