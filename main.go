package main

import (
	"os"

	di "github.com/ppdx999/kyopro/internal/dependency_injection"
)

func main() {
	os.Exit(int(di.Cmd().Run(os.Args[1:])))
}
