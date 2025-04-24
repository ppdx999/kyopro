package session

import (
	"github.com/ppdx999/kyopro/internal/domain/model"
)

type SessionLoader interface {
	LoadSession() (model.SessionSecret, error)
}

type SessionLoaderImpl struct {
	sessionPath    SessionPath
	existFile      ExistFile
	readSecretFile ReadSecretFile
}

func NewSessionLoaderImpl(
	sessionPath SessionPath,
	existFile ExistFile,
	readSecretFile ReadSecretFile,
) *SessionLoaderImpl {
	return &SessionLoaderImpl{
		sessionPath:    sessionPath,
		existFile:      existFile,
		readSecretFile: readSecretFile,
	}
}

func (l *SessionLoaderImpl) LoadSession() (model.SessionSecret, error) {
	path, err := l.sessionPath.SessionPath()
	if err != nil {
		return "", err
	}
	if !l.existFile.ExistFile(path) {
		return "", nil
	}

	data, err := l.readSecretFile.ReadSecretFile(path)
	if err != nil {
		return "", err
	}
	return model.SessionSecret(data), nil
}
