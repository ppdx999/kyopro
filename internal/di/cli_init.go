package di

import init_ "github.com/ppdx999/kyopro/internal/cli/init"

func InitializeInitCli() *init_.InitCli {
	consoleConsole := InitializeConsole()
	initService := InitializeInitService()
	return init_.NewInitCli(consoleConsole, initService)
}
