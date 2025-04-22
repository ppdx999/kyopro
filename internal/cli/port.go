package cli

type MsgSender interface {
	SendMsg(msg string)
}
