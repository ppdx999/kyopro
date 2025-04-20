package model

type ProblemId string

type Problem struct {
	ID ProblemId
}

func NewProblem(id string) *Problem {
	return &Problem{
		ID: ProblemId(id),
	}
}
