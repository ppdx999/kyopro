package di

import application_service "github.com/ppdx999/kyopro/internal/application/service"

func ApplicationServiceTest() application_service.Tester {
	return application_service.NewTester(
		UserPipeline(),
		TestCaseCurrentLoader(),
		LanguageDetector(),
		LanguageTestCaseRunner(),
		MsgSender(),
	)
}
