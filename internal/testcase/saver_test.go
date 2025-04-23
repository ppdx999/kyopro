package testcase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/testcase"
)

type MockGetWd struct {
	wd  string
	err error
}

func (m *MockGetWd) GetWd() (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.wd, nil
}

type MockDirMaker struct {
	makedDirs []string
	err       error
}

func (m *MockDirMaker) MakePublicDir(dir string) error {
	if m.err != nil {
		return m.err
	}
	m.makedDirs = append(m.makedDirs, dir)
	return nil
}

type MockPublicFileWriter struct {
	writtenFiles map[string][]byte
	err          error
}

func (m *MockPublicFileWriter) WritePublicFile(path string, data []byte) error {
	if m.err != nil {
		return m.err
	}
	m.writtenFiles[path] = data
	return nil
}

func TestSaveTestCase(t *testing.T) {
	tests := []struct {
		name               string
		tc                 *model.TestCase
		getWd              string
		getWdErr           error
		makePublicDirErr   error
		writtenFiles       map[string][]byte
		writePublicFileErr error
		wantErr            bool
	}{
		{
			name: "success",
			tc: &model.TestCase{
				ID:    "1",
				Input: []byte("input 1.in"),
				Want:  []byte("output 1.out"),
			},
			getWd: "/tmp",
			writtenFiles: map[string][]byte{
				"/tmp/test/1.in":  []byte("input 1.in"),
				"/tmp/test/1.out": []byte("output 1.out"),
			},
		},
		{
			name:         "getWd error",
			tc:           &model.TestCase{ID: "1"},
			getWdErr:     fmt.Errorf("getWd error"),
			writtenFiles: make(map[string][]byte),
			wantErr:      true,
		},
		{
			name:             "makePublicDir error",
			tc:               &model.TestCase{ID: "1"},
			getWd:            "/tmp",
			makePublicDirErr: fmt.Errorf("makePublicDir error"),
			writtenFiles:     make(map[string][]byte),
			wantErr:          true,
		},
		{
			name:               "writePublicFile error",
			tc:                 &model.TestCase{ID: "1"},
			getWd:              "/tmp",
			writtenFiles:       make(map[string][]byte),
			writePublicFileErr: fmt.Errorf("writePublicFile error"),
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGetWd := &MockGetWd{
				wd:  tt.getWd,
				err: tt.getWdErr,
			}
			mockDirMaker := &MockDirMaker{
				err: tt.makePublicDirErr,
			}
			mockPublicFileWriter := &MockPublicFileWriter{
				writtenFiles: make(map[string][]byte),
				err:          tt.writePublicFileErr,
			}
			saver := testcase.NewTestCaseFsSaver(mockGetWd, mockDirMaker, mockPublicFileWriter)

			err := saver.SaveTestCase(tt.tc)

			if (err != nil) != tt.wantErr {
				t.Errorf("TestCaseFsSaver.Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.writtenFiles, mockPublicFileWriter.writtenFiles) {
				t.Errorf("TestCaseFsSaver.Save() writtenFiles = %v, want %v", mockPublicFileWriter.writtenFiles, tt.writtenFiles)
			}
		})
	}
}
