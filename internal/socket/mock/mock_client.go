// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_socket is a generated GoMock package.
package mock_socket

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSocket is a mock of Socket interface.
type MockSocket struct {
	ctrl     *gomock.Controller
	recorder *MockSocketMockRecorder
}

// MockSocketMockRecorder is the mock recorder for MockSocket.
type MockSocketMockRecorder struct {
	mock *MockSocket
}

// NewMockSocket creates a new mock instance.
func NewMockSocket(ctrl *gomock.Controller) *MockSocket {
	mock := &MockSocket{ctrl: ctrl}
	mock.recorder = &MockSocketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocket) EXPECT() *MockSocketMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSocket) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSocketMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSocket)(nil).Read), arg0)
}

// Write mocks base method.
func (m *MockSocket) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSocketMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSocket)(nil).Write), arg0)
}
