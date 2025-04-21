package di

import (
	"github.com/ppdx999/kyopro/internal/session"
)

func InitializeSessionPath() session.SessionPath {
	var InitializeHome = func() session.Home {
		return InitializeUserHome()
	}

	home := InitializeHome()
	return session.NewSessionPath(home)
}

func InitializeAskSessionImpl() *session.AskSessionImpl {
	var InitializeUserInput = func() session.UserInput {
		return InitializeUserInputFromConsole()
	}

	userInput := InitializeUserInput()
	return session.NewAskSessionImpl(userInput)
}

func InitializeLoadSessionImpl() *session.LoadSessionImpl {
	var InitializeReadSecretFile = func() session.ReadSecretFile {
		return InitializeFsImpl()
	}

	sessionPath := InitializeSessionPath()
	readSecretFile := InitializeReadSecretFile()
	return session.NewLoadSessionImpl(sessionPath, readSecretFile)
}

func InitializeSaveSessionImpl() *session.SaveSessionImpl {
	var InitializeMakePublicDir = func() session.MakePublicDir {
		return InitializeFsImpl()
	}

	var InitializeWriteSecretFile = func() session.WriteSecretFile {
		return InitializeFsImpl()
	}

	sessionPath := InitializeSessionPath()
	makePublicDir := InitializeMakePublicDir()
	writeSecretFile := InitializeWriteSecretFile()

	return session.NewSaveSessionImpl(sessionPath, makePublicDir, writeSecretFile)
}
