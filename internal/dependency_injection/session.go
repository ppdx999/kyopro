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

	return session.NewSessionAskerImpl(UserInput())
}

func SessionLoader() session.SessionLoader {
	var ReadSecretFile = func() session.ReadSecretFile {
		return FsImpl()
	}

	var ExitFile = func() session.ExistFile {
		return FsImpl()
	}

	return session.NewSessionLoaderImpl(
		SessionPath(),
		ExitFile(),
		ReadSecretFile(),
	)
}

func SessionSaver() session.SessionSaver {
	var MakePublicDir = func() session.MakePublicDir {
		return FsImpl()
	}

	var WriteSecretFile = func() session.WriteSecretFile {
		return FsImpl()
	}

	return session.NewSessionSaverImpl(
		SessionPath(),
		MakePublicDir(),
		WriteSecretFile(),
	)
}
