// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/dot/digest (interfaces: GrandpaState)

// Package digest is a generated GoMock package.
package digest

import (
	reflect "reflect"

	types "github.com/ChainSafe/gossamer/dot/types"
	scale "github.com/ChainSafe/gossamer/pkg/scale"
	gomock "github.com/golang/mock/gomock"
)

// MockGrandpaState is a mock of GrandpaState interface.
type MockGrandpaState struct {
	ctrl     *gomock.Controller
	recorder *MockGrandpaStateMockRecorder
}

// MockGrandpaStateMockRecorder is the mock recorder for MockGrandpaState.
type MockGrandpaStateMockRecorder struct {
	mock *MockGrandpaState
}

// NewMockGrandpaState creates a new mock instance.
func NewMockGrandpaState(ctrl *gomock.Controller) *MockGrandpaState {
	mock := &MockGrandpaState{ctrl: ctrl}
	mock.recorder = &MockGrandpaStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGrandpaState) EXPECT() *MockGrandpaStateMockRecorder {
	return m.recorder
}

// ApplyForcedChanges mocks base method.
func (m *MockGrandpaState) ApplyForcedChanges(arg0 *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyForcedChanges", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyForcedChanges indicates an expected call of ApplyForcedChanges.
func (mr *MockGrandpaStateMockRecorder) ApplyForcedChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyForcedChanges", reflect.TypeOf((*MockGrandpaState)(nil).ApplyForcedChanges), arg0)
}

// ApplyScheduledChanges mocks base method.
func (m *MockGrandpaState) ApplyScheduledChanges(arg0 *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyScheduledChanges", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyScheduledChanges indicates an expected call of ApplyScheduledChanges.
func (mr *MockGrandpaStateMockRecorder) ApplyScheduledChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyScheduledChanges", reflect.TypeOf((*MockGrandpaState)(nil).ApplyScheduledChanges), arg0)
}

// GetCurrentSetID mocks base method.
func (m *MockGrandpaState) GetCurrentSetID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSetID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSetID indicates an expected call of GetCurrentSetID.
func (mr *MockGrandpaStateMockRecorder) GetCurrentSetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSetID", reflect.TypeOf((*MockGrandpaState)(nil).GetCurrentSetID))
}

// HandleGRANDPADigest mocks base method.
func (m *MockGrandpaState) HandleGRANDPADigest(arg0 *types.Header, arg1 scale.VaryingDataType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleGRANDPADigest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleGRANDPADigest indicates an expected call of HandleGRANDPADigest.
func (mr *MockGrandpaStateMockRecorder) HandleGRANDPADigest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGRANDPADigest", reflect.TypeOf((*MockGrandpaState)(nil).HandleGRANDPADigest), arg0, arg1)
}

// IncrementSetID mocks base method.
func (m *MockGrandpaState) IncrementSetID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementSetID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrementSetID indicates an expected call of IncrementSetID.
func (mr *MockGrandpaStateMockRecorder) IncrementSetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementSetID", reflect.TypeOf((*MockGrandpaState)(nil).IncrementSetID))
}

// SetNextChange mocks base method.
func (m *MockGrandpaState) SetNextChange(arg0 []types.GrandpaVoter, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNextChange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNextChange indicates an expected call of SetNextChange.
func (mr *MockGrandpaStateMockRecorder) SetNextChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNextChange", reflect.TypeOf((*MockGrandpaState)(nil).SetNextChange), arg0, arg1)
}

// SetNextPause mocks base method.
func (m *MockGrandpaState) SetNextPause(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNextPause", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNextPause indicates an expected call of SetNextPause.
func (mr *MockGrandpaStateMockRecorder) SetNextPause(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNextPause", reflect.TypeOf((*MockGrandpaState)(nil).SetNextPause), arg0)
}

// SetNextResume mocks base method.
func (m *MockGrandpaState) SetNextResume(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNextResume", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNextResume indicates an expected call of SetNextResume.
func (mr *MockGrandpaStateMockRecorder) SetNextResume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNextResume", reflect.TypeOf((*MockGrandpaState)(nil).SetNextResume), arg0)
}