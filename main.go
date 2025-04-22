package main

import (
	"os"

	"github.com/ppdx999/kyopro/internal/di"
)

func main() {
	cmd := di.InitializeCmd()
	os.Exit(int(cmd.Run(os.Args[1:])))
}
