package test_mocks

// This one's part generated, part hand-written - it's stateful, so we need
// to give it at least a little implementation-y goodness or else things won't
// work right

import (
	"github.com/NGTOne/warren/conn"

	"github.com/golang/mock/gomock"
	"reflect"
)

type TestConnection struct {
	ctrl     *gomock.Controller
	recorder *TestConnectionMockRecorder
	callback func(conn.Message)

	msg conn.Message
}

type TestConnectionMockRecorder struct {
	mock *TestConnection
}

func NewTestConnection(
	msg conn.Message,
	ctrl *gomock.Controller,
) *TestConnection {
	mock := &TestConnection{ctrl: ctrl, msg: msg}
	mock.recorder = &TestConnectionMockRecorder{mock}
	return mock
}

func (m *TestConnection) SetNewMsgCallback(f func(conn.Message)) {
	m.callback = f
}

// Borrowed from gomock's generated code
func (m *TestConnection) EXPECT() *TestConnectionMockRecorder {
	return m.recorder
}

// AckMsg mocks base method
func (m *TestConnection) AckMsg(arg0 conn.Message) error {
	ret := m.ctrl.Call(m, "AckMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AckMsg indicates an expected call of AckMsg
func (mr *TestConnectionMockRecorder) AckMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckMsg", reflect.TypeOf((*TestConnection)(nil).AckMsg), arg0)
}

// Listen mocks base method
func (m *TestConnection) Listen(f func(msg conn.Message)) {
	f(m.msg)
}

// NackMsg mocks base method
func (m *TestConnection) NackMsg(arg0 conn.Message) error {
	ret := m.ctrl.Call(m, "NackMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NackMsg indicates an expected call of NackMsg
func (mr *TestConnectionMockRecorder) NackMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NackMsg", reflect.TypeOf((*TestConnection)(nil).NackMsg), arg0)
}

// SendResponse mocks base method
func (m *TestConnection) SendResponse(arg0, arg1 conn.Message) error {
	ret := m.ctrl.Call(m, "SendResponse", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResponse indicates an expected call of SendResponse
func (mr *TestConnectionMockRecorder) SendResponse(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResponse", reflect.TypeOf((*TestConnection)(nil).SendResponse), arg0, arg1)
}
