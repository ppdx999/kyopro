package session

import (
	"github.com/ppdx999/kyopro/internal/model"
)

type SessionLoaderImpl struct {
	sessionPath    SessionPath
	readSecretFile ReadSecretFile
}

func NewSessionLoaderImpl(
	sessionPath SessionPath,
	readSecretFile ReadSecretFile,
) *SessionLoaderImpl {
	return &SessionLoaderImpl{
		sessionPath:    sessionPath,
		readSecretFile: readSecretFile,
	}
}

func (l *SessionLoaderImpl) LoadSession() (model.SessionSecret, error) {
	path, err := l.sessionPath.SessionPath()
	if err != nil {
		return "", err
	}
	data, err := l.readSecretFile.ReadSecretFile(path)
	if err != nil {
		return "", err
	}
	return model.SessionSecret(data), nil
}
