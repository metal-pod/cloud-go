// Code generated by mockery v2.7.4. DO NOT EDIT.

package volume

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	volume "github.com/fi-ts/cloud-go/api/client/volume"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// ClusterInfo provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ClusterInfo(params *volume.ClusterInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...volume.ClientOption) (*volume.ClusterInfoOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *volume.ClusterInfoOK
	if rf, ok := ret.Get(0).(func(*volume.ClusterInfoParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) *volume.ClusterInfoOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.ClusterInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*volume.ClusterInfoParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolume provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteVolume(params *volume.DeleteVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...volume.ClientOption) (*volume.DeleteVolumeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *volume.DeleteVolumeOK
	if rf, ok := ret.Get(0).(func(*volume.DeleteVolumeParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) *volume.DeleteVolumeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.DeleteVolumeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*volume.DeleteVolumeParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindVolumes provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindVolumes(params *volume.FindVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...volume.ClientOption) (*volume.FindVolumesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *volume.FindVolumesOK
	if rf, ok := ret.Get(0).(func(*volume.FindVolumesParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) *volume.FindVolumesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.FindVolumesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*volume.FindVolumesParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVolume provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetVolume(params *volume.GetVolumeParams, authInfo runtime.ClientAuthInfoWriter, opts ...volume.ClientOption) (*volume.GetVolumeOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *volume.GetVolumeOK
	if rf, ok := ret.Get(0).(func(*volume.GetVolumeParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) *volume.GetVolumeOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.GetVolumeOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*volume.GetVolumeParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVolumes provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListVolumes(params *volume.ListVolumesParams, authInfo runtime.ClientAuthInfoWriter, opts ...volume.ClientOption) (*volume.ListVolumesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *volume.ListVolumesOK
	if rf, ok := ret.Get(0).(func(*volume.ListVolumesParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) *volume.ListVolumesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*volume.ListVolumesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*volume.ListVolumesParams, runtime.ClientAuthInfoWriter, ...volume.ClientOption) error); ok {
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
