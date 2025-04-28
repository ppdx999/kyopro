package di

import "github.com/ppdx999/kyopro/internal/infrastructure/browser"

func Browser() *browser.Browser {
	return browser.NewBrowser()
}
