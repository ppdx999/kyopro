package testutil

import "reflect"

type MockMsgSender struct {
	sendMsgs []string
}

func (m *MockMsgSender) SendMsg(msg string) {
	m.sendMsgs = append(m.sendMsgs, msg)
}

func (m *MockMsgSender) CheckMsgs(msgs []string) bool {
	if len(m.sendMsgs) != len(msgs) {
		return false
	}
	if !reflect.DeepEqual(m.sendMsgs, msgs) {
		return false
	}
	return true
}

func (m *MockMsgSender) Clear() {
	m.sendMsgs = nil
}

func (m *MockMsgSender) GetSendMsgs() []string {
	return m.sendMsgs
}
