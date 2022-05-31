// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	clienttenant "github.com/fi-ts/cloud-go/api/client/tenant"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"

	testing "testing"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// FindTenants provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindTenants(params *clienttenant.FindTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clienttenant.ClientOption) (*clienttenant.FindTenantsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clienttenant.FindTenantsOK
	if rf, ok := ret.Get(0).(func(*clienttenant.FindTenantsParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) *clienttenant.FindTenantsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clienttenant.FindTenantsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clienttenant.FindTenantsParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTenant provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetTenant(params *clienttenant.GetTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...clienttenant.ClientOption) (*clienttenant.GetTenantOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clienttenant.GetTenantOK
	if rf, ok := ret.Get(0).(func(*clienttenant.GetTenantParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) *clienttenant.GetTenantOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clienttenant.GetTenantOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clienttenant.GetTenantParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTenants provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListTenants(params *clienttenant.ListTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clienttenant.ClientOption) (*clienttenant.ListTenantsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clienttenant.ListTenantsOK
	if rf, ok := ret.Get(0).(func(*clienttenant.ListTenantsParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) *clienttenant.ListTenantsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clienttenant.ListTenantsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clienttenant.ListTenantsParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateTenant provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateTenant(params *clienttenant.UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...clienttenant.ClientOption) (*clienttenant.UpdateTenantOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clienttenant.UpdateTenantOK
	if rf, ok := ret.Get(0).(func(*clienttenant.UpdateTenantParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) *clienttenant.UpdateTenantOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clienttenant.UpdateTenantOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clienttenant.UpdateTenantParams, runtime.ClientAuthInfoWriter, ...clienttenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientService creates a new instance of ClientService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t testing.TB) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
