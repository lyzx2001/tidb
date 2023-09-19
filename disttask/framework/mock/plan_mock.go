// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/disttask/framework/planner (interfaces: LogicalPlan,PipelineSpec)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	planner "github.com/pingcap/tidb/disttask/framework/planner"
	gomock "go.uber.org/mock/gomock"
)

// MockLogicalPlan is a mock of LogicalPlan interface.
type MockLogicalPlan struct {
	ctrl     *gomock.Controller
	recorder *MockLogicalPlanMockRecorder
}

// MockLogicalPlanMockRecorder is the mock recorder for MockLogicalPlan.
type MockLogicalPlanMockRecorder struct {
	mock *MockLogicalPlan
}

// NewMockLogicalPlan creates a new mock instance.
func NewMockLogicalPlan(ctrl *gomock.Controller) *MockLogicalPlan {
	mock := &MockLogicalPlan{ctrl: ctrl}
	mock.recorder = &MockLogicalPlanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogicalPlan) EXPECT() *MockLogicalPlanMockRecorder {
	return m.recorder
}

// FromTaskMeta mocks base method.
func (m *MockLogicalPlan) FromTaskMeta(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromTaskMeta", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FromTaskMeta indicates an expected call of FromTaskMeta.
func (mr *MockLogicalPlanMockRecorder) FromTaskMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromTaskMeta", reflect.TypeOf((*MockLogicalPlan)(nil).FromTaskMeta), arg0)
}

// ToPhysicalPlan mocks base method.
func (m *MockLogicalPlan) ToPhysicalPlan(arg0 planner.PlanCtx) (*planner.PhysicalPlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToPhysicalPlan", arg0)
	ret0, _ := ret[0].(*planner.PhysicalPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToPhysicalPlan indicates an expected call of ToPhysicalPlan.
func (mr *MockLogicalPlanMockRecorder) ToPhysicalPlan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToPhysicalPlan", reflect.TypeOf((*MockLogicalPlan)(nil).ToPhysicalPlan), arg0)
}

// ToTaskMeta mocks base method.
func (m *MockLogicalPlan) ToTaskMeta() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToTaskMeta")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToTaskMeta indicates an expected call of ToTaskMeta.
func (mr *MockLogicalPlanMockRecorder) ToTaskMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToTaskMeta", reflect.TypeOf((*MockLogicalPlan)(nil).ToTaskMeta))
}

// MockPipelineSpec is a mock of PipelineSpec interface.
type MockPipelineSpec struct {
	ctrl     *gomock.Controller
	recorder *MockPipelineSpecMockRecorder
}

// MockPipelineSpecMockRecorder is the mock recorder for MockPipelineSpec.
type MockPipelineSpecMockRecorder struct {
	mock *MockPipelineSpec
}

// NewMockPipelineSpec creates a new mock instance.
func NewMockPipelineSpec(ctrl *gomock.Controller) *MockPipelineSpec {
	mock := &MockPipelineSpec{ctrl: ctrl}
	mock.recorder = &MockPipelineSpecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelineSpec) EXPECT() *MockPipelineSpecMockRecorder {
	return m.recorder
}

// ToSubtaskMeta mocks base method.
func (m *MockPipelineSpec) ToSubtaskMeta(arg0 planner.PlanCtx) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSubtaskMeta", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToSubtaskMeta indicates an expected call of ToSubtaskMeta.
func (mr *MockPipelineSpecMockRecorder) ToSubtaskMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSubtaskMeta", reflect.TypeOf((*MockPipelineSpec)(nil).ToSubtaskMeta), arg0)
}
