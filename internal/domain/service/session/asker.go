package session

import "github.com/ppdx999/kyopro/internal/domain/model"

type sessionAsker struct {
	userInput UserInput
}

func (s *sessionAsker) AskSession() (model.SessionSecret, error) {
	input, err := s.userInput.UserInput()
	if err != nil {
		return "", err
	}
	return model.SessionSecret(input), nil
}
