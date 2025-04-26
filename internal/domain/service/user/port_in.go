package user

type LoginChecker interface {
	LoginCheck() (bool, error)
}

type MsgSender interface {
	SendMsg(string)
}

type UserInput interface {
	UserInput() (string, error)
}

type Home interface {
	Home() (string, error)
}
