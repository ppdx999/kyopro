package testcase

import (
	"errors"
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type testCaseSaver struct {
	wd         WdGetter
	dirMaker   PublicDirMaker
	fileWriter PublicFileWriter
}

func (t *testCaseSaver) SaveTestCase(ts *model.TestCase) error {
	cwd, err := t.wd.GetWd()
	if err != nil {
		return err
	}
	testdir := filepath.Join(cwd, "test")
	if err := t.dirMaker.MakePublicDir(testdir); err != nil {
		return err
	}

	inputPath := filepath.Join(testdir, string(ts.ID)+".in")
	wantPath := filepath.Join(testdir, string(ts.ID)+".out")

	if err = errors.Join(
		t.fileWriter.WritePublicFile(inputPath, ts.Input),
		t.fileWriter.WritePublicFile(wantPath, ts.Want),
	); err != nil {
		return err
	}

	return nil
}
