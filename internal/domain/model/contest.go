package model

type ContestId string

type Contest struct {
	ID       ContestId
	Problems []Problem
}

func NewContest(id string) *Contest {
	if id == "" {
		panic("contest id is empty")
	}
	return &Contest{
		ID: ContestId(id),
	}
}
