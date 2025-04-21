package init

import (
	"github.com/ppdx999/kyopro/internal/model"
)

type InitService interface {
	Init(c model.ContestId) error
}

/*
InitServiceImplは問題の一覧を取得して、それぞれの問題のディレクトリを作成します。
*/
type InitServiceImpl struct {
	GetProblemIds  GetProblemIds
	MakeProblemDir MakeProblemDir
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

func NewInitServiceImpl(
	GetProblemIds GetProblemIds,
	MakeProblemDir MakeProblemDir,
) *InitServiceImpl {
	return &InitServiceImpl{
		GetProblemIds:  GetProblemIds,
		MakeProblemDir: MakeProblemDir,
	}
}
