package application_service

import (
	"bytes"
	"fmt"

	"github.com/ppdx999/kyopro/internal/domain/service/language"
	"github.com/ppdx999/kyopro/internal/domain/service/testcase"
	"github.com/ppdx999/kyopro/internal/domain/service/user"
)

type tester struct {
	userPipeline           user.Pipeline
	testCasesCurrentLoader testcase.TestCaseCurrentLoader
	languageDetector       language.LanguageDetector
	languageTestcaseRunner language.LanguageTestCaseRunner
	msgSender              user.MsgSender
}

func NewTester(
	userPipeline user.Pipeline,
	testCasesCurrentLoader testcase.TestCaseCurrentLoader,
	languageDetector language.LanguageDetector,
	languageTestcaseRunner language.LanguageTestCaseRunner,
	msgSender user.MsgSender,
) *tester {
	return &tester{
		userPipeline:           userPipeline,
		testCasesCurrentLoader: testCasesCurrentLoader,
		languageDetector:       languageDetector,
		languageTestcaseRunner: languageTestcaseRunner,
		msgSender:              msgSender,
	}
}

func (t *tester) Test() error {
	ts, err := t.testCasesCurrentLoader.LoadCurrentTestCases()
	if err != nil {
		return err
	}
	lang, err := t.languageDetector.DetectLanguage()
	if err != nil {
		return err
	}

	pipeline := t.userPipeline.Pipeline()
	pipeline.Inflow = bytes.NewReader(nil)
	if err := lang.Build(pipeline); err != nil {
		return err
	}

	for _, tc := range ts {
		got, errMsg, err := t.languageTestcaseRunner.Run(lang, tc)
		if err != nil {
			return err
		}
		if len(errMsg) > 0 {
			t.msgSender.SendMsg(string(errMsg))
		}

		buf := bytes.NewBuffer(nil)
		if bytes.Equal(got, tc.Want) {
			fmt.Fprintf(buf, "✅ Test %s passed\n", tc.ID)
		} else {
			fmt.Fprintf(buf, "❌ Test %s failed\n", tc.ID)
			fmt.Fprintln(buf, "Input:")
			fmt.Fprintf(buf, "%s\n", tc.Input)
			fmt.Fprintln(buf, "Want:")
			fmt.Fprintf(buf, "%s\n", tc.Want)
			fmt.Fprintln(buf, "Got:")
			fmt.Fprintf(buf, "%s\n", got)
		}
	}

	pipeline = t.userPipeline.Pipeline()
	pipeline.Inflow = bytes.NewReader(nil)
	if err := lang.Clean(pipeline); err != nil {
		return err
	}

	return nil
}
