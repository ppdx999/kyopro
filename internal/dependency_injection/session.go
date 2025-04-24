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

func SessionAskerImpl() *session.SessionAskerImpl {
	var UserInput = func() session.UserInput {
		return UserInputFromConsole()
	}

	return session.NewSessionAskerImpl(UserInput())
}

func SessionLoaderImpl() *session.SessionLoaderImpl {
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

func SessionSaverImpl() *session.SessionSaverImpl {
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
