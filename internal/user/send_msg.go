package user

import "fmt"

type SendMsgByConsole struct{}

func (s *SendMsgByConsole) SendMsg(msg string) {
	fmt.Println(msg)
}
