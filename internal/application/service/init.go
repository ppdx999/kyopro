package application_service

import (
	"github.com/ppdx999/kyopro/internal/domain/model"
	"github.com/ppdx999/kyopro/internal/domain/service/problem"
)

/*
initerは問題の一覧を取得して、それぞれの問題のディレクトリを作成します。
*/
type initer struct {
	getter problem.ProblemIdsGetter
	maker  problem.ProblemDirMaker
}

func NewIniter(g problem.ProblemIdsGetter, m problem.ProblemDirMaker) *initer {
	return &initer{getter: g, maker: m}
}

func (s *initer) Init(c model.ContestId) error {
	ids, err := s.getter.GetProblemIds(c)
	if err != nil {
		return err
	}
	for _, id := range ids {
		err := s.maker.MakeProblemDir(c, id)
		if err != nil {
			return err
		}
	}
	return nil
}
