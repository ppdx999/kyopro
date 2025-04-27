package session

import "github.com/ppdx999/kyopro/internal/domain/model"

type SessionAsker interface {
	AskSession() (model.SessionSecret, error)
}

func NewSessionAsker(userInput UserInput) SessionAsker {
	return &sessionAsker{
		userInput: userInput,
	}
}

type SessionLoader interface {
	LoadSession() (model.SessionSecret, error)
}

func NewSessionLoader(
	sessionPath SessionPath,
	existFile ExistFile,
	readSecretFile ReadSecretFile,
) SessionLoader {
	return &sessionLoader{
		sessionPath:    sessionPath,
		existFile:      existFile,
		readSecretFile: readSecretFile,
	}
}

type SessionSaver interface {
	SaveSession(model.SessionSecret) error
}

func NewSessionSaver(
	sessionPath SessionPath,
	makePublicDir MakePublicDir,
	writeSecretFile WriteSecretFile,
) SessionSaver {
	return &sessionSaver{
		sessionPath:     sessionPath,
		makePublicDir:   makePublicDir,
		writeSecretFile: writeSecretFile,
	}
}
