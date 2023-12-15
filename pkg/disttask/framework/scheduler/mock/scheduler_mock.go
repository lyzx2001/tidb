// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/pkg/disttask/framework/scheduler (interfaces: Extension)
//
// Generated by this command:
//
//	mockgen -package mock github.com/pingcap/tidb/pkg/disttask/framework/scheduler Extension
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	proto "github.com/pingcap/tidb/pkg/disttask/framework/proto"
	scheduler "github.com/pingcap/tidb/pkg/disttask/framework/scheduler"
	infosync "github.com/pingcap/tidb/pkg/domain/infosync"
	gomock "go.uber.org/mock/gomock"
)

// MockExtension is a mock of Extension interface.
type MockExtension struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionMockRecorder
}

// MockExtensionMockRecorder is the mock recorder for MockExtension.
type MockExtensionMockRecorder struct {
	mock *MockExtension
}

// NewMockExtension creates a new mock instance.
func NewMockExtension(ctrl *gomock.Controller) *MockExtension {
	mock := &MockExtension{ctrl: ctrl}
	mock.recorder = &MockExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension) EXPECT() *MockExtensionMockRecorder {
	return m.recorder
}

// GetEligibleInstances mocks base method.
func (m *MockExtension) GetEligibleInstances(arg0 context.Context, arg1 *proto.Task) ([]*infosync.ServerInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEligibleInstances", arg0, arg1)
	ret0, _ := ret[0].([]*infosync.ServerInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEligibleInstances indicates an expected call of GetEligibleInstances.
func (mr *MockExtensionMockRecorder) GetEligibleInstances(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEligibleInstances", reflect.TypeOf((*MockExtension)(nil).GetEligibleInstances), arg0, arg1)
}

// GetNextStep mocks base method.
func (m *MockExtension) GetNextStep(arg0 *proto.Task) proto.Step {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextStep", arg0)
	ret0, _ := ret[0].(proto.Step)
	return ret0
}

// GetNextStep indicates an expected call of GetNextStep.
func (mr *MockExtensionMockRecorder) GetNextStep(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextStep", reflect.TypeOf((*MockExtension)(nil).GetNextStep), arg0)
}

// IsRetryableErr mocks base method.
func (m *MockExtension) IsRetryableErr(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRetryableErr", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRetryableErr indicates an expected call of IsRetryableErr.
func (mr *MockExtensionMockRecorder) IsRetryableErr(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRetryableErr", reflect.TypeOf((*MockExtension)(nil).IsRetryableErr), arg0)
}

// OnDone mocks base method.
func (m *MockExtension) OnDone(arg0 context.Context, arg1 scheduler.TaskHandle, arg2 *proto.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnDone", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnDone indicates an expected call of OnDone.
func (mr *MockExtensionMockRecorder) OnDone(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDone", reflect.TypeOf((*MockExtension)(nil).OnDone), arg0, arg1, arg2)
}

// OnNextSubtasksBatch mocks base method.
func (m *MockExtension) OnNextSubtasksBatch(arg0 context.Context, arg1 scheduler.TaskHandle, arg2 *proto.Task, arg3 []*infosync.ServerInfo, arg4 proto.Step) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnNextSubtasksBatch", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnNextSubtasksBatch indicates an expected call of OnNextSubtasksBatch.
func (mr *MockExtensionMockRecorder) OnNextSubtasksBatch(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNextSubtasksBatch", reflect.TypeOf((*MockExtension)(nil).OnNextSubtasksBatch), arg0, arg1, arg2, arg3, arg4)
}

// OnTick mocks base method.
func (m *MockExtension) OnTick(arg0 context.Context, arg1 *proto.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnTick", arg0, arg1)
}

// OnTick indicates an expected call of OnTick.
func (mr *MockExtensionMockRecorder) OnTick(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnTick", reflect.TypeOf((*MockExtension)(nil).OnTick), arg0, arg1)
}