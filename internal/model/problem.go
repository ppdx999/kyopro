package model

type ProblemId string

type Problem struct {
	ID      ProblemId
	Contest *Contest
}

func NewProblem(id string) *Problem {
	if id == "" {
		panic("problem id is empty")
	}

	return &Problem{
		ID: ProblemId(id),
	}
}
