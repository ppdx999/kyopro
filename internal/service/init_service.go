package service

import (
	"github.com/ppdx999/kyopro/internal/model"
	"github.com/ppdx999/kyopro/internal/service/helper"
)

/*
InitServiceは問題の一覧を取得して、それぞれの問題のディレクトリを作成します。
*/
type InitService interface {
	Init(c model.ContestId) error
}

type InitServiceImpl struct {
	GetProblemIds  helper.GetProblemIds
	MakeProblemDir helper.MakeProblemDir
}

func (s *InitServiceImpl) Init(c model.ContestId) error {
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
