package main

import (
	"os"

	"github.com/ppdx999/kyopro/internal/di"
)

func main() {
	os.Exit(int(di.Cmd().Run(os.Args[1:])))
}
