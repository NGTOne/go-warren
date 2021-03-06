// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NGTOne/warren/service (interfaces: ErrHandler)

// Package test_mocks is a generated GoMock package.
package test_mocks

import (
	conn "github.com/NGTOne/warren/conn"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockErrHandler is a mock of ErrHandler interface
type MockErrHandler struct {
	ctrl     *gomock.Controller
	recorder *MockErrHandlerMockRecorder
}

// MockErrHandlerMockRecorder is the mock recorder for MockErrHandler
type MockErrHandlerMockRecorder struct {
	mock *MockErrHandler
}

// NewMockErrHandler creates a new mock instance
func NewMockErrHandler(ctrl *gomock.Controller) *MockErrHandler {
	mock := &MockErrHandler{ctrl: ctrl}
	mock.recorder = &MockErrHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockErrHandler) EXPECT() *MockErrHandlerMockRecorder {
	return m.recorder
}

// ProcessErr mocks base method
func (m *MockErrHandler) ProcessErr(arg0 conn.Message, arg1 error) error {
	ret := m.ctrl.Call(m, "ProcessErr", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessErr indicates an expected call of ProcessErr
func (mr *MockErrHandlerMockRecorder) ProcessErr(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessErr", reflect.TypeOf((*MockErrHandler)(nil).ProcessErr), arg0, arg1)
}
