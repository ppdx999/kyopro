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

func InitializeSessionAskerImpl() *session.SessionAskerImpl {
	var InitializeUserInput = func() session.UserInput {
		return InitializeUserInputFromConsole()
	}

	userInput := InitializeUserInput()
	return session.NewSessionAskerImpl(userInput)
}

func InitializeSessionLoaderImpl() *session.SessionLoaderImpl {
	var InitializeReadSecretFile = func() session.ReadSecretFile {
		return InitializeFsImpl()
	}

	sessionPath := InitializeSessionPath()
	readSecretFile := InitializeReadSecretFile()
	return session.NewSessionLoaderImpl(sessionPath, readSecretFile)
}

func InitializeSessionSaverImpl() *session.SessionSaverImpl {
	var InitializeMakePublicDir = func() session.MakePublicDir {
		return InitializeFsImpl()
	}

	var InitializeWriteSecretFile = func() session.WriteSecretFile {
		return InitializeFsImpl()
	}

	sessionPath := InitializeSessionPath()
	makePublicDir := InitializeMakePublicDir()
	writeSecretFile := InitializeWriteSecretFile()

	return session.NewSessionSaverImpl(sessionPath, makePublicDir, writeSecretFile)
}
