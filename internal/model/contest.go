package model

type ContestId string

type Contest struct {
	ID ContestId
}

func NewContest(id string) *Contest {
	return &Contest{
		ID: ContestId(id),
	}
}
