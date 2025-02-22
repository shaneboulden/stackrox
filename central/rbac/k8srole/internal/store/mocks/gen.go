// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/rbac/k8srole/internal/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteRole mocks base method
func (m *MockStore) DeleteRole(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole
func (mr *MockStoreMockRecorder) DeleteRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockStore)(nil).DeleteRole), arg0)
}

// GetRole mocks base method
func (m *MockStore) GetRole(arg0 string) (*storage.K8SRole, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(*storage.K8SRole)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRole indicates an expected call of GetRole
func (mr *MockStoreMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockStore)(nil).GetRole), arg0)
}

// GetRoles mocks base method
func (m *MockStore) GetRoles(arg0 []string) ([]*storage.K8SRole, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", arg0)
	ret0, _ := ret[0].([]*storage.K8SRole)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRoles indicates an expected call of GetRoles
func (mr *MockStoreMockRecorder) GetRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockStore)(nil).GetRoles), arg0)
}

// ListRoles mocks base method
func (m *MockStore) ListRoles() ([]*storage.K8SRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles")
	ret0, _ := ret[0].([]*storage.K8SRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles
func (mr *MockStoreMockRecorder) ListRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockStore)(nil).ListRoles))
}

// UpsertRole mocks base method
func (m *MockStore) UpsertRole(arg0 *storage.K8SRole) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRole", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRole indicates an expected call of UpsertRole
func (mr *MockStoreMockRecorder) UpsertRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRole", reflect.TypeOf((*MockStore)(nil).UpsertRole), arg0)
}
