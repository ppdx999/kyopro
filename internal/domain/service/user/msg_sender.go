package user

import "fmt"

type ConsoleMsgSender struct{}

func (s *ConsoleMsgSender) SendMsg(msg string) {
	fmt.Println(msg)
}
