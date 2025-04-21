package user

import "os"

type UserHome struct{}

func (u *UserHome) Home() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}
