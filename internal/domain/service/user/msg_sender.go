package user

import "fmt"

type MsgSender interface {
	SendMsg(string)
}

type ConsoleMsgSender struct{}

func (s *ConsoleMsgSender) SendMsg(msg string) {
	fmt.Println(msg)
}
