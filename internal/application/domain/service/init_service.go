package service

import (
	"github.com/ppdx999/kyopro/internal/application/domain/model"
	"github.com/ppdx999/kyopro/internal/application/port/out"
)

type InitService struct {
	getProblemIds  out.GetProblemIdsPort
	makeProblemDir out.MakeProblemDirPort
}

func (s *InitService) Init(c model.ContestId) error {
	ids, err := s.getProblemIds.GetProblemIds(c)
	if err != nil {
		return err
	}
	for _, id := range ids {
		err := s.makeProblemDir.MakeProblemDir(c, id)
		if err != nil {
			return err
		}
	}
	return nil
}
