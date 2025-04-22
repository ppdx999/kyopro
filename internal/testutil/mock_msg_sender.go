package testutil

type MockMsgSender struct {
	sendMsgs []string
}

func (m *MockMsgSender) SendMsg(msg string) {
	m.sendMsgs = append(m.sendMsgs, msg)
}
