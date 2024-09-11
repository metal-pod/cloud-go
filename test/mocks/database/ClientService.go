// Code generated by mockery v2.45.1. DO NOT EDIT.

package database

import (
	clientdatabase "github.com/fi-ts/cloud-go/api/client/database"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AcceptPostgresRestore provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AcceptPostgresRestore(params *clientdatabase.AcceptPostgresRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.AcceptPostgresRestoreOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AcceptPostgresRestore")
	}

	var r0 *clientdatabase.AcceptPostgresRestoreOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.AcceptPostgresRestoreParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.AcceptPostgresRestoreOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.AcceptPostgresRestoreParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.AcceptPostgresRestoreOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.AcceptPostgresRestoreOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.AcceptPostgresRestoreParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptPostgresRestoreDeprecated provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AcceptPostgresRestoreDeprecated(params *clientdatabase.AcceptPostgresRestoreDeprecatedParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.AcceptPostgresRestoreDeprecatedOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AcceptPostgresRestoreDeprecated")
	}

	var r0 *clientdatabase.AcceptPostgresRestoreDeprecatedOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.AcceptPostgresRestoreDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.AcceptPostgresRestoreDeprecatedOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.AcceptPostgresRestoreDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.AcceptPostgresRestoreDeprecatedOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.AcceptPostgresRestoreDeprecatedOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.AcceptPostgresRestoreDeprecatedParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePostgres(params *clientdatabase.CreatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePostgres")
	}

	var r0 *clientdatabase.CreatePostgresCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.CreatePostgresCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.CreatePostgresCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.CreatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePostgresBackupConfig(params *clientdatabase.CreatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresBackupConfigCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePostgresBackupConfig")
	}

	var r0 *clientdatabase.CreatePostgresBackupConfigCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresBackupConfigCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.CreatePostgresBackupConfigCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.CreatePostgresBackupConfigCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.CreatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePostgresStandby provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePostgresStandby(params *clientdatabase.CreatePostgresStandbyParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresStandbyCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePostgresStandby")
	}

	var r0 *clientdatabase.CreatePostgresStandbyCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresStandbyParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.CreatePostgresStandbyCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.CreatePostgresStandbyParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.CreatePostgresStandbyCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.CreatePostgresStandbyCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.CreatePostgresStandbyParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeletePostgres(params *clientdatabase.DeletePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.DeletePostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePostgres")
	}

	var r0 *clientdatabase.DeletePostgresOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.DeletePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.DeletePostgresOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.DeletePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.DeletePostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.DeletePostgresOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.DeletePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeletePostgresBackupConfig(params *clientdatabase.DeletePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.DeletePostgresBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePostgresBackupConfig")
	}

	var r0 *clientdatabase.DeletePostgresBackupConfigOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.DeletePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.DeletePostgresBackupConfigOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.DeletePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.DeletePostgresBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.DeletePostgresBackupConfigOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.DeletePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindPostgres(params *clientdatabase.FindPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.FindPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindPostgres")
	}

	var r0 *clientdatabase.FindPostgresOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.FindPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.FindPostgresOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.FindPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.FindPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.FindPostgresOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.FindPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetBackupConfig(params *clientdatabase.GetBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBackupConfig")
	}

	var r0 *clientdatabase.GetBackupConfigOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetBackupConfigOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetBackupConfigOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgres(params *clientdatabase.GetPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgres")
	}

	var r0 *clientdatabase.GetPostgresOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetPostgresOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresBackups provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresBackups(params *clientdatabase.GetPostgresBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresBackupsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgresBackups")
	}

	var r0 *clientdatabase.GetPostgresBackupsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresBackupsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresBackupsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresBackupsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetPostgresBackupsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetPostgresBackupsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetPostgresBackupsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresPartitions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresPartitions(params *clientdatabase.GetPostgresPartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresPartitionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgresPartitions")
	}

	var r0 *clientdatabase.GetPostgresPartitionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresPartitionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresPartitionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresPartitionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetPostgresPartitionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetPostgresPartitionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetPostgresPartitionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresSecrets provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresSecrets(params *clientdatabase.GetPostgresSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresSecretsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgresSecrets")
	}

	var r0 *clientdatabase.GetPostgresSecretsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresSecretsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresSecretsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresSecretsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetPostgresSecretsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetPostgresSecretsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetPostgresSecretsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresVersions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresVersions(params *clientdatabase.GetPostgresVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresVersionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPostgresVersions")
	}

	var r0 *clientdatabase.GetPostgresVersionsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresVersionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.GetPostgresVersionsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.GetPostgresVersionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.GetPostgresVersionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.GetPostgresVersionsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.GetPostgresVersionsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListPostgres(params *clientdatabase.ListPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.ListPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPostgres")
	}

	var r0 *clientdatabase.ListPostgresOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.ListPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.ListPostgresOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.ListPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.ListPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.ListPostgresOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.ListPostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPostgresBackupConfigs provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListPostgresBackupConfigs(params *clientdatabase.ListPostgresBackupConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.ListPostgresBackupConfigsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPostgresBackupConfigs")
	}

	var r0 *clientdatabase.ListPostgresBackupConfigsOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.ListPostgresBackupConfigsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.ListPostgresBackupConfigsOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.ListPostgresBackupConfigsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.ListPostgresBackupConfigsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.ListPostgresBackupConfigsOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.ListPostgresBackupConfigsParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestorePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) RestorePostgres(params *clientdatabase.RestorePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.RestorePostgresCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RestorePostgres")
	}

	var r0 *clientdatabase.RestorePostgresCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.RestorePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.RestorePostgresCreated, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.RestorePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.RestorePostgresCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.RestorePostgresCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.RestorePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
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

// UpdatePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdatePostgres(params *clientdatabase.UpdatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.UpdatePostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePostgres")
	}

	var r0 *clientdatabase.UpdatePostgresOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.UpdatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.UpdatePostgresOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.UpdatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.UpdatePostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.UpdatePostgresOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.UpdatePostgresParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdatePostgresBackupConfig(params *clientdatabase.UpdatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientdatabase.ClientOption) (*clientdatabase.UpdatePostgresBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePostgresBackupConfig")
	}

	var r0 *clientdatabase.UpdatePostgresBackupConfigOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientdatabase.UpdatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) (*clientdatabase.UpdatePostgresBackupConfigOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientdatabase.UpdatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) *clientdatabase.UpdatePostgresBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientdatabase.UpdatePostgresBackupConfigOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientdatabase.UpdatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...clientdatabase.ClientOption) error); ok {
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
