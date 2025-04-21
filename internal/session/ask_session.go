package session

import "github.com/ppdx999/kyopro/internal/model"

type AskSessionImpl struct {
	userInput UserInput
}

func NewAskSessionImpl(userInput UserInput) *AskSessionImpl {
	return &AskSessionImpl{
		userInput: userInput,
	}
}

func (s *AskSessionImpl) AskSession() (model.SessionSecret, error) {
	input, err := s.userInput.UserInput()
	if err != nil {
		return "", err
	}
	return model.SessionSecret(input), nil
}
