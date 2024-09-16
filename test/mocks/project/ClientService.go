// Code generated by mockery v2.45.1. DO NOT EDIT.

package project

import (
	clientproject "github.com/fi-ts/cloud-go/api/client/project"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateMachineReservation provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateMachineReservation(params *clientproject.CreateMachineReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.CreateMachineReservationCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMachineReservation")
	}

	var r0 *clientproject.CreateMachineReservationCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.CreateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.CreateMachineReservationCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.CreateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.CreateMachineReservationCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.CreateMachineReservationCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.CreateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateProject(params *clientproject.CreateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.CreateProjectCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 *clientproject.CreateProjectCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.CreateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.CreateProjectCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.CreateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.CreateProjectCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.CreateProjectCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.CreateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMachineReservation provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteMachineReservation(params *clientproject.DeleteMachineReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.DeleteMachineReservationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMachineReservation")
	}

	var r0 *clientproject.DeleteMachineReservationOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.DeleteMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.DeleteMachineReservationOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.DeleteMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.DeleteMachineReservationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.DeleteMachineReservationOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.DeleteMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteProject(params *clientproject.DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.DeleteProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProject")
	}

	var r0 *clientproject.DeleteProjectOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.DeleteProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.DeleteProjectOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.DeleteProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.DeleteProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.DeleteProjectOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.DeleteProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindProject(params *clientproject.FindProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.FindProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindProject")
	}

	var r0 *clientproject.FindProjectOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.FindProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.FindProjectOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.FindProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.FindProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.FindProjectOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.FindProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProjects provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindProjects(params *clientproject.FindProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.FindProjectsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindProjects")
	}

	var r0 *clientproject.FindProjectsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.FindProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.FindProjectsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.FindProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.FindProjectsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.FindProjectsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.FindProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMachineReservations provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListMachineReservations(params *clientproject.ListMachineReservationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.ListMachineReservationsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMachineReservations")
	}

	var r0 *clientproject.ListMachineReservationsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.ListMachineReservationsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.ListMachineReservationsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.ListMachineReservationsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.ListMachineReservationsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.ListMachineReservationsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.ListMachineReservationsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListProjects(params *clientproject.ListProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.ListProjectsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjects")
	}

	var r0 *clientproject.ListProjectsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.ListProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.ListProjectsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.ListProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.ListProjectsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.ListProjectsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.ListProjectsParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
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

// UpdateMachineReservation provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateMachineReservation(params *clientproject.UpdateMachineReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.UpdateMachineReservationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMachineReservation")
	}

	var r0 *clientproject.UpdateMachineReservationOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.UpdateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.UpdateMachineReservationOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.UpdateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.UpdateMachineReservationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.UpdateMachineReservationOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.UpdateMachineReservationParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateProject(params *clientproject.UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientproject.ClientOption) (*clientproject.UpdateProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProject")
	}

	var r0 *clientproject.UpdateProjectOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientproject.UpdateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) (*clientproject.UpdateProjectOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientproject.UpdateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) *clientproject.UpdateProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientproject.UpdateProjectOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientproject.UpdateProjectParams, runtime.ClientAuthInfoWriter, ...clientproject.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
