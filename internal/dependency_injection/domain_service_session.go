package di

import (
	"github.com/ppdx999/kyopro/internal/domain/service/session"
)

func SessionPath() session.SessionPath {
	var Home = func() session.Home {
		return UserHome()
	}

	return session.NewSessionPath(Home())
}

func SessionAsker() session.SessionAsker {
	var UserInput = func() session.UserInput {
		return UserInputFromConsole()
	}

	return session.NewSessionAsker(UserInput())
}

func SessionLoader() session.SessionLoader {
	var ReadSecretFile = func() session.ReadSecretFile {
		return FileSystem()
	}

	var ExitFile = func() session.ExistFile {
		return FileSystem()
	}

	return session.NewSessionLoader(
		SessionPath(),
		ExitFile(),
		ReadSecretFile(),
	)
}

func SessionSaver() session.SessionSaver {
	var MakePublicDir = func() session.MakePublicDir {
		return FileSystem()
	}

	var WriteSecretFile = func() session.WriteSecretFile {
		return FileSystem()
	}

	return session.NewSessionSaver(
		SessionPath(),
		MakePublicDir(),
		WriteSecretFile(),
	)
}
