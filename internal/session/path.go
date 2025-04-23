package session

import "path/filepath"

type SessionPath interface {
	SessionPath() (string, error)
}

type SessionPathImpl struct {
	home Home
}

func NewSessionPath(home Home) *SessionPathImpl {
	return &SessionPathImpl{
		home: home,
	}
}

func (s *SessionPathImpl) SessionPath() (string, error) {
	home, err := s.home.Home()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".local", "share")
	return filepath.Join(dir, "kyopro", "session.txt"), nil
}
