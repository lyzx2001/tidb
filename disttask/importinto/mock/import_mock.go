// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/disttask/importinto (interfaces: MiniTaskExecutor)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMiniTaskExecutor is a mock of MiniTaskExecutor interface.
type MockMiniTaskExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockMiniTaskExecutorMockRecorder
}

// MockMiniTaskExecutorMockRecorder is the mock recorder for MockMiniTaskExecutor.
type MockMiniTaskExecutorMockRecorder struct {
	mock *MockMiniTaskExecutor
}

// NewMockMiniTaskExecutor creates a new mock instance.
func NewMockMiniTaskExecutor(ctrl *gomock.Controller) *MockMiniTaskExecutor {
	mock := &MockMiniTaskExecutor{ctrl: ctrl}
	mock.recorder = &MockMiniTaskExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiniTaskExecutor) EXPECT() *MockMiniTaskExecutorMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockMiniTaskExecutor) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockMiniTaskExecutorMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMiniTaskExecutor)(nil).Run), arg0)
}
