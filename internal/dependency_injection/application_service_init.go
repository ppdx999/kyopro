package di

import (
	application_service "github.com/ppdx999/kyopro/internal/application/service"
)

func ApplicationServiceInit() application_service.Initer {
	return application_service.NewIniter(
		ProblemIdsGetter(),
		ProblemDirMaker(),
	)
}
