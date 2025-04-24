package init

import (
	"github.com/ppdx999/kyopro/internal/domain/model"
)

type InitService interface {
	Init(c model.ContestId) error
}

/*
InitServiceImplは問題の一覧を取得して、それぞれの問題のディレクトリを作成します。
*/
type InitServiceImpl struct {
	getter GetProblemIds
	maker  ProblemDirMaker
}

func NewInitServiceImpl(g GetProblemIds, m ProblemDirMaker) *InitServiceImpl {
	return &InitServiceImpl{getter: g, maker: m}
}

func (s *InitServiceImpl) Init(c model.ContestId) error {
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
