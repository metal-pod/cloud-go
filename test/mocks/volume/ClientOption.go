// Code generated by mockery v2.12.2. DO NOT EDIT.

package volume

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ClientOption is an autogenerated mock type for the ClientOption type
type ClientOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ClientOption) Execute(_a0 *runtime.ClientOperation) {
	_m.Called(_a0)
}

// NewClientOption creates a new instance of ClientOption. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientOption(t testing.TB) *ClientOption {
	mock := &ClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
