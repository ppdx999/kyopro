package application_service

import (
	"github.com/ppdx999/kyopro/internal/domain/service/language"
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
)

type submiter struct {
	languageDetector language.LanguageDetector
	problemLoader    problem.CurrentProblemLoader
	clipboardWriter  ClipboardWriter
	submitPageOpner  SubmitPageOpener
}

func NewSubmiter(
	languageDetector language.LanguageDetector,
	problemLoader problem.CurrentProblemLoader,
	clipboardWriter ClipboardWriter,
	submitPageOpner SubmitPageOpener,
) Submiter {
	return &submiter{
		languageDetector: languageDetector,
		problemLoader:    problemLoader,
		clipboardWriter:  clipboardWriter,
		submitPageOpner:  submitPageOpner,
	}
}

func (s *submiter) Submit() error {
	lang, err := s.languageDetector.DetectLanguage()
	if err != nil {
		return err
	}
	sourceCode := lang.SourceCode()
	s.clipboardWriter.WriteClipboard(sourceCode.Code)

	problem, err := s.problemLoader.LoadCurrentProblem()
	if err != nil {
		return err
	}

	return s.submitPageOpner.OpenSubmitPage(problem.Contest.ID, problem.ID)
}
