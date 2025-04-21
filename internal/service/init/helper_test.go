package init_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	service "github.com/ppdx999/kyopro/internal/service/init"
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

			service := service.GetWorkspaceImpl{
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

type MockMakePublicDir struct {
	createdDirs []string
	err         error
}

func NewMockMakePublicDir() *MockMakePublicDir {
	return &MockMakePublicDir{}
}

func (m *MockMakePublicDir) MakePublicDir(path string) error {
	if m.err != nil {
		return m.err
	}
	m.createdDirs = append(m.createdDirs, path)
	return m.err
}

type MockGetWorkspace struct {
	workspace model.Workspace
	err       error
}

func NewMockGetWorkspace() *MockGetWorkspace {
	return &MockGetWorkspace{}
}

func (m *MockGetWorkspace) GetWorkspace() (*model.Workspace, error) {
	return &m.workspace, m.err
}

func TestMakeProblemDir(t *testing.T) {
	tests := []struct {
		name             string
		contest          model.ContestId
		problem          model.ProblemId
		workspace        model.Workspace
		makePublicDirErr error
		getWorkspaceErr  error
		wantErr          bool
		createdDirs      []string
	}{
		{
			name:             "正常系",
			contest:          model.ContestId("abc123"),
			problem:          model.ProblemId("1"),
			workspace:        model.Workspace{Path: "/path/to/workspace"},
			makePublicDirErr: nil,
			getWorkspaceErr:  nil,
			wantErr:          false,
			createdDirs: []string{
				"/path/to/workspace/abc123/1",
			},
		},
		{
			name:             "MakePublicDirエラー",
			contest:          model.ContestId("abc123"),
			problem:          model.ProblemId("1"),
			workspace:        model.Workspace{Path: "/path/to/workspace"},
			makePublicDirErr: errors.New("make public dir error"),
			getWorkspaceErr:  nil,
			wantErr:          true,
			createdDirs:      nil,
		},
		{
			name:             "GetWorkspaceエラー",
			contest:          model.ContestId("abc123"),
			problem:          model.ProblemId("1"),
			workspace:        model.Workspace{Path: "/path/to/workspace"},
			makePublicDirErr: nil,
			getWorkspaceErr:  errors.New("get workspace error"),
			wantErr:          true,
			createdDirs:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makePublicDir := NewMockMakePublicDir()
			makePublicDir.err = tt.makePublicDirErr
			getWorkspace := NewMockGetWorkspace()
			getWorkspace.workspace = tt.workspace
			getWorkspace.err = tt.getWorkspaceErr

			service := &service.MakeProblemDirImpl{
				MakePublicDir: makePublicDir,
				GetWorkspace:  getWorkspace,
			}

			err := service.MakeProblemDir(tt.contest, tt.problem)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeProblemDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(makePublicDir.createdDirs, tt.createdDirs) {
				t.Errorf("MakeProblemDir() createdDirs = %v, want %v", makePublicDir.createdDirs, tt.createdDirs)
				return
			}
		})
	}
}
