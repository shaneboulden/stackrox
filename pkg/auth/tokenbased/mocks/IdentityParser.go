// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import metadata "google.golang.org/grpc/metadata"
import mock "github.com/stretchr/testify/mock"
import tokenbased "github.com/stackrox/rox/pkg/auth/tokenbased"

// IdentityParser is an autogenerated mock type for the IdentityParser type
type IdentityParser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: md, roleMapper
func (_m *IdentityParser) Parse(md metadata.MD, roleMapper tokenbased.RoleMapper) (tokenbased.Identity, error) {
	ret := _m.Called(md, roleMapper)

	var r0 tokenbased.Identity
	if rf, ok := ret.Get(0).(func(metadata.MD, tokenbased.RoleMapper) tokenbased.Identity); ok {
		r0 = rf(md, roleMapper)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tokenbased.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metadata.MD, tokenbased.RoleMapper) error); ok {
		r1 = rf(md, roleMapper)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
