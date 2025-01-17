// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChainSafe/gossamer/lib/runtime (interfaces: Memory)

// Package runtime is a generated GoMock package.
package runtime

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMemory is a mock of Memory interface.
type MockMemory struct {
	ctrl     *gomock.Controller
	recorder *MockMemoryMockRecorder
}

// MockMemoryMockRecorder is the mock recorder for MockMemory.
type MockMemoryMockRecorder struct {
	mock *MockMemory
}

// NewMockMemory creates a new mock instance.
func NewMockMemory(ctrl *gomock.Controller) *MockMemory {
	mock := &MockMemory{ctrl: ctrl}
	mock.recorder = &MockMemoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemory) EXPECT() *MockMemoryMockRecorder {
	return m.recorder
}

// Data mocks base method.
func (m *MockMemory) Data() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockMemoryMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockMemory)(nil).Data))
}

// Grow mocks base method.
func (m *MockMemory) Grow(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Grow indicates an expected call of Grow.
func (mr *MockMemoryMockRecorder) Grow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grow", reflect.TypeOf((*MockMemory)(nil).Grow), arg0)
}

// Length mocks base method.
func (m *MockMemory) Length() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockMemoryMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMemory)(nil).Length))
}
