package application_service

import "github.com/ppdx999/kyopro/internal/domain/model"

type Downloader interface {
	Download() error
}

type Initer interface {
	Init(c model.ContestId) error
}

type Loginer interface {
	Login() error
}
