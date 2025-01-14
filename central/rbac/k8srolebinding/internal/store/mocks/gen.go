// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/rbac/k8srolebinding/internal/store (interfaces: Store)

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

// DeleteRoleBinding mocks base method
func (m *MockStore) DeleteRoleBinding(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoleBinding", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoleBinding indicates an expected call of DeleteRoleBinding
func (mr *MockStoreMockRecorder) DeleteRoleBinding(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoleBinding", reflect.TypeOf((*MockStore)(nil).DeleteRoleBinding), arg0)
}

// GetRoleBinding mocks base method
func (m *MockStore) GetRoleBinding(arg0 string) (*storage.K8SRoleBinding, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleBinding", arg0)
	ret0, _ := ret[0].(*storage.K8SRoleBinding)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRoleBinding indicates an expected call of GetRoleBinding
func (mr *MockStoreMockRecorder) GetRoleBinding(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleBinding", reflect.TypeOf((*MockStore)(nil).GetRoleBinding), arg0)
}

// GetRoleBindings mocks base method
func (m *MockStore) GetRoleBindings(arg0 []string) ([]*storage.K8SRoleBinding, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleBindings", arg0)
	ret0, _ := ret[0].([]*storage.K8SRoleBinding)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRoleBindings indicates an expected call of GetRoleBindings
func (mr *MockStoreMockRecorder) GetRoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleBindings", reflect.TypeOf((*MockStore)(nil).GetRoleBindings), arg0)
}

// ListRoleBindings mocks base method
func (m *MockStore) ListRoleBindings() ([]*storage.K8SRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleBindings")
	ret0, _ := ret[0].([]*storage.K8SRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleBindings indicates an expected call of ListRoleBindings
func (mr *MockStoreMockRecorder) ListRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleBindings", reflect.TypeOf((*MockStore)(nil).ListRoleBindings))
}

// UpsertRoleBinding mocks base method
func (m *MockStore) UpsertRoleBinding(arg0 *storage.K8SRoleBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRoleBinding", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRoleBinding indicates an expected call of UpsertRoleBinding
func (mr *MockStoreMockRecorder) UpsertRoleBinding(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRoleBinding", reflect.TypeOf((*MockStore)(nil).UpsertRoleBinding), arg0)
}
