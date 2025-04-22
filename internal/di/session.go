package di

import (
	"github.com/ppdx999/kyopro/internal/session"
)

func SessionPath() session.SessionPath {
	var Home = func() session.Home {
		return UserHome()
	}

	home := Home()
	return session.NewSessionPath(home)
}

func SessionAskerImpl() *session.SessionAskerImpl {
	var UserInput = func() session.UserInput {
		return UserInputFromConsole()
	}

	userInput := UserInput()
	return session.NewSessionAskerImpl(userInput)
}

func SessionLoaderImpl() *session.SessionLoaderImpl {
	var ReadSecretFile = func() session.ReadSecretFile {
		return FsImpl()
	}

	var ExitFile = func() session.ExistFile {
		return FsImpl()
	}

	sessionPath := SessionPath()
	exitFile := ExitFile()
	readSecretFile := ReadSecretFile()
	return session.NewSessionLoaderImpl(sessionPath, exitFile, readSecretFile)
}

func SessionSaverImpl() *session.SessionSaverImpl {
	var MakePublicDir = func() session.MakePublicDir {
		return FsImpl()
	}

	var WriteSecretFile = func() session.WriteSecretFile {
		return FsImpl()
	}

	sessionPath := SessionPath()
	makePublicDir := MakePublicDir()
	writeSecretFile := WriteSecretFile()

	return session.NewSessionSaverImpl(sessionPath, makePublicDir, writeSecretFile)
}
