package session

import "github.com/ppdx999/kyopro/internal/domain/model"

type SessionAsker interface {
	AskSession() (model.SessionSecret, error)
}

type SessionAskerImpl struct {
	userInput UserInput
}

func NewSessionAskerImpl(userInput UserInput) *SessionAskerImpl {
	return &SessionAskerImpl{
		userInput: userInput,
	}
}

func (s *SessionAskerImpl) AskSession() (model.SessionSecret, error) {
	input, err := s.userInput.UserInput()
	if err != nil {
		return "", err
	}
	return model.SessionSecret(input), nil
}
