package user

import "fmt"

type UserInputFromConsole struct{}

func (u UserInputFromConsole) UserInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}
