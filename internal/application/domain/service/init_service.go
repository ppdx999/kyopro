package service

import (
	"github.com/ppdx999/kyopro/internal/application/domain/model"
	"github.com/ppdx999/kyopro/internal/application/port/out"
)

type InitService struct {
	GetProblemIds  out.GetProblemIdsPort
	MakeProblemDir out.MakeProblemDirPort
}

func (s *InitService) Init(c model.ContestId) error {
	ids, err := s.GetProblemIds.GetProblemIds(c)
	if err != nil {
		return err
	}
	for _, id := range ids {
		err := s.MakeProblemDir.MakeProblemDir(c, id)
		if err != nil {
			return err
		}
	}
	return nil
}
