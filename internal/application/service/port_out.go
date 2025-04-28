package application_service

import "github.com/ppdx999/kyopro/internal/domain/model"

type ClipboardWriter interface {
	WriteClipboard([]byte) error
}

type SubmitPageOpener interface {
	OpenSubmitPage(
		c model.ContestId,
		p model.ProblemId,
	) error
}
