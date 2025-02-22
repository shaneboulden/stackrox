// Code generated by MockGen. DO NOT EDIT.
// Source: handler_impl.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCompressor is a mock of Compressor interface
type MockCompressor struct {
	ctrl     *gomock.Controller
	recorder *MockCompressorMockRecorder
}

// MockCompressorMockRecorder is the mock recorder for MockCompressor
type MockCompressorMockRecorder struct {
	mock *MockCompressor
}

// NewMockCompressor creates a new mock instance
func NewMockCompressor(ctrl *gomock.Controller) *MockCompressor {
	mock := &MockCompressor{ctrl: ctrl}
	mock.recorder = &MockCompressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompressor) EXPECT() *MockCompressorMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockCompressor) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockCompressorMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCompressor)(nil).Write), p)
}

// Close mocks base method
func (m *MockCompressor) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCompressorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCompressor)(nil).Close))
}

// Bytes mocks base method
func (m *MockCompressor) Bytes() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Bytes indicates an expected call of Bytes
func (mr *MockCompressorMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockCompressor)(nil).Bytes))
}
