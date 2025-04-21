package session

import (
	"github.com/ppdx999/kyopro/internal/model"
)

type LoadSessionImpl struct {
	sessionPath    SessionPath
	readSecretFile ReadSecretFile
}

func NewLoadSessionImpl(
	sessionPath SessionPath,
	readSecretFile ReadSecretFile,
) *LoadSessionImpl {
	return &LoadSessionImpl{
		sessionPath:    sessionPath,
		readSecretFile: readSecretFile,
	}
}

func (l *LoadSessionImpl) LoadSession() (model.SessionSecret, error) {
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
