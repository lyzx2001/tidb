// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/br/pkg/lightning/backend/encode (interfaces: Encoder,EncodingBuilder,Rows,Row)
//
// Generated by this command:
//
//	mockgen -package mock github.com/pingcap/tidb/br/pkg/lightning/backend/encode Encoder,EncodingBuilder,Rows,Row
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	encode "github.com/pingcap/tidb/br/pkg/lightning/backend/encode"
	verification "github.com/pingcap/tidb/br/pkg/lightning/verification"
	types "github.com/pingcap/tidb/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockEncoder is a mock of Encoder interface.
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEncoder) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEncoderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEncoder)(nil).Close))
}

// Encode mocks base method.
func (m *MockEncoder) Encode(arg0 []types.Datum, arg1 int64, arg2 []int, arg3 int64) (encode.Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(encode.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderMockRecorder) Encode(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0, arg1, arg2, arg3)
}

// MockEncodingBuilder is a mock of EncodingBuilder interface.
type MockEncodingBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockEncodingBuilderMockRecorder
}

// MockEncodingBuilderMockRecorder is the mock recorder for MockEncodingBuilder.
type MockEncodingBuilderMockRecorder struct {
	mock *MockEncodingBuilder
}

// NewMockEncodingBuilder creates a new mock instance.
func NewMockEncodingBuilder(ctrl *gomock.Controller) *MockEncodingBuilder {
	mock := &MockEncodingBuilder{ctrl: ctrl}
	mock.recorder = &MockEncodingBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncodingBuilder) EXPECT() *MockEncodingBuilderMockRecorder {
	return m.recorder
}

// MakeEmptyRows mocks base method.
func (m *MockEncodingBuilder) MakeEmptyRows() encode.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeEmptyRows")
	ret0, _ := ret[0].(encode.Rows)
	return ret0
}

// MakeEmptyRows indicates an expected call of MakeEmptyRows.
func (mr *MockEncodingBuilderMockRecorder) MakeEmptyRows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeEmptyRows", reflect.TypeOf((*MockEncodingBuilder)(nil).MakeEmptyRows))
}

// NewEncoder mocks base method.
func (m *MockEncodingBuilder) NewEncoder(arg0 context.Context, arg1 *encode.EncodingConfig) (encode.Encoder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEncoder", arg0, arg1)
	ret0, _ := ret[0].(encode.Encoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewEncoder indicates an expected call of NewEncoder.
func (mr *MockEncodingBuilderMockRecorder) NewEncoder(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEncoder", reflect.TypeOf((*MockEncodingBuilder)(nil).NewEncoder), arg0, arg1)
}

// MockRows is a mock of Rows interface.
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows.
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance.
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockRows) Clear() encode.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(encode.Rows)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockRowsMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockRows)(nil).Clear))
}

// MockRow is a mock of Row interface.
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow.
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance.
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// ClassifyAndAppend mocks base method.
func (m *MockRow) ClassifyAndAppend(arg0 *encode.Rows, arg1 *verification.KVChecksum, arg2 *encode.Rows, arg3 *verification.KVChecksum) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClassifyAndAppend", arg0, arg1, arg2, arg3)
}

// ClassifyAndAppend indicates an expected call of ClassifyAndAppend.
func (mr *MockRowMockRecorder) ClassifyAndAppend(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassifyAndAppend", reflect.TypeOf((*MockRow)(nil).ClassifyAndAppend), arg0, arg1, arg2, arg3)
}

// Size mocks base method.
func (m *MockRow) Size() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockRowMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockRow)(nil).Size))
}
