package session

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/model"
)

type SaveSessionImpl struct {
	sessionPath     SessionPath
	makePublicDir   MakePublicDir
	writeSecretFile WriteSecretFile
}

func NewSaveSessionImpl(
	sessionPath SessionPath,
	makePublicDir MakePublicDir,
	writeSecretFile WriteSecretFile,
) *SaveSessionImpl {
	return &SaveSessionImpl{
		sessionPath:     sessionPath,
		makePublicDir:   makePublicDir,
		writeSecretFile: writeSecretFile,
	}
}

func (s *SaveSessionImpl) SaveSession(session model.SessionSecret) error {
	path, err := s.sessionPath.SessionPath()
	if err != nil {
		return err
	}
	if err := s.makePublicDir.MakePublicDir(filepath.Dir(path)); err != nil {
		return err
	}
	if err := s.writeSecretFile.WriteSecretFile(path, []byte(session)); err != nil {
		return err
	}
	return nil
}
