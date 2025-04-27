package session

import (
	"path/filepath"

	"github.com/ppdx999/kyopro/internal/domain/model"
)

type sessionSaver struct {
	sessionPath     SessionPath
	makePublicDir   MakePublicDir
	writeSecretFile WriteSecretFile
}

func (s *sessionSaver) SaveSession(session model.SessionSecret) error {
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
