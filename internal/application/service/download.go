package application_service

import (
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
	"github.com/ppdx999/kyopro/internal/domain/service/testcase"
)

/*
downloaderは問題のテストケースをダウンロードします。
*/
type downloader struct {
	loader problem.CurrentProblemLoader
	getter testcase.TestCasesGetter
	saver  testcase.TestCaseSaver
}

func NewDownloader(
	loader problem.CurrentProblemLoader,
	getter testcase.TestCasesGetter,
	saver testcase.TestCaseSaver,
) *downloader {
	return &downloader{
		loader: loader,
		getter: getter,
		saver:  saver,
	}
}

func (d *downloader) Download() error {
	p, err := d.loader.LoadCurrentProblem()
	if err != nil {
		return err
	}

	testCases, err := d.getter.GetTestCases(
		p.Contest.ID,
		p.ID,
	)
	if err != nil {
		return err
	}

	for _, t := range testCases {
		if err := d.saver.SaveTestCase(t); err != nil {
			return err
		}
	}

	return nil
}
