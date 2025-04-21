package di

import (
	"github.com/ppdx999/kyopro/internal/console"
)

func InitializeConsole() console.Console {
	return console.NewConsoleImpl()
}
