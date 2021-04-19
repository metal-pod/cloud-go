// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new database API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for database API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePostgres(params *CreatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostgresCreated, error)

	CreatePostgresBackupConfig(params *CreatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostgresBackupConfigCreated, error)

	DeletePostgres(params *DeletePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostgresOK, error)

	DeletePostgresBackupConfig(params *DeletePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostgresBackupConfigOK, error)

	FindPostgres(params *FindPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPostgresOK, error)

	GetBackupConfig(params *GetBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupConfigOK, error)

	GetPostgres(params *GetPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresOK, error)

	GetPostgresBackups(params *GetPostgresBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresBackupsOK, error)

	GetPostgresPartitions(params *GetPostgresPartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresPartitionsOK, error)

	GetPostgresSecrets(params *GetPostgresSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresSecretsOK, error)

	GetPostgresVersions(params *GetPostgresVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresVersionsOK, error)

	ListPostgres(params *ListPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostgresOK, error)

	ListPostgresBackupConfigs(params *ListPostgresBackupConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostgresBackupConfigsOK, error)

	UpdatePostgres(params *UpdatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostgresOK, error)

	UpdatePostgresBackupConfig(params *UpdatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostgresBackupConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePostgres creates a postgres if the given ID already exists a conflict is returned
*/
func (a *Client) CreatePostgres(params *CreatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostgresCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPostgres",
		Method:             "PUT",
		PathPattern:        "/v1/database/postgres",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostgresCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreatePostgresBackupConfig creates a postgres backup for the given projectid
*/
func (a *Client) CreatePostgresBackupConfig(params *CreatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreatePostgresBackupConfigCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostgresBackupConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPostgresBackupConfig",
		Method:             "PUT",
		PathPattern:        "/v1/database/postgres/backup-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePostgresBackupConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostgresBackupConfigCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePostgresBackupConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePostgres deletes an postgres and returns the deleted entity
*/
func (a *Client) DeletePostgres(params *DeletePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostgresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePostgres",
		Method:             "DELETE",
		PathPattern:        "/v1/database/postgres/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePostgresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeletePostgresBackupConfig deletes a postgres backup for the given projectid
*/
func (a *Client) DeletePostgresBackupConfig(params *DeletePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePostgresBackupConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePostgresBackupConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePostgresBackupConfig",
		Method:             "DELETE",
		PathPattern:        "/v1/database/postgres/backup-config/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePostgresBackupConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePostgresBackupConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeletePostgresBackupConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindPostgres finds postgres databases by multiple criteria
*/
func (a *Client) FindPostgres(params *FindPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPostgresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPostgres",
		Method:             "POST",
		PathPattern:        "/v1/database/postgres/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindPostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindPostgresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindPostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBackupConfig gets postgres backups
*/
func (a *Client) GetBackupConfig(params *GetBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBackupConfig",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/backup-config/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackupConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBackupConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBackupConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPostgres gets postgres by id
*/
func (a *Client) GetPostgres(params *GetPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostgres",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostgresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPostgresBackups gets postgres stored backups by id
*/
func (a *Client) GetPostgresBackups(params *GetPostgresBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresBackupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresBackupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostgresBackups",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/{id}/backups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPostgresBackupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostgresBackupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresBackupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPostgresPartitions gets postgres partitions supported
*/
func (a *Client) GetPostgresPartitions(params *GetPostgresPartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresPartitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresPartitionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostgresPartitions",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/partitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPostgresPartitionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostgresPartitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresPartitionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPostgresSecrets gets postgres secrets by id
*/
func (a *Client) GetPostgresSecrets(params *GetPostgresSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresSecretsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostgresSecrets",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/{id}/secrets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPostgresSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostgresSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresSecretsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetPostgresVersions gets postgres versions supported
*/
func (a *Client) GetPostgresVersions(params *GetPostgresVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPostgresVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostgresVersions",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPostgresVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPostgresVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListPostgres gets all postgres databases
*/
func (a *Client) ListPostgres(params *ListPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostgresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPostgres",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostgresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListPostgresBackupConfigs gets all postgres backups
*/
func (a *Client) ListPostgresBackupConfigs(params *ListPostgresBackupConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPostgresBackupConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostgresBackupConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPostgresBackupConfigs",
		Method:             "GET",
		PathPattern:        "/v1/database/postgres/backup-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPostgresBackupConfigsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostgresBackupConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListPostgresBackupConfigsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdatePostgres updates a postgres if the postgres was changed since this one was read a conflict is returned
*/
func (a *Client) UpdatePostgres(params *UpdatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostgresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePostgresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePostgres",
		Method:             "POST",
		PathPattern:        "/v1/database/postgres",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePostgresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePostgresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePostgresDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdatePostgresBackupConfig updates a postgres backup for the given projectid
*/
func (a *Client) UpdatePostgresBackupConfig(params *UpdatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePostgresBackupConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePostgresBackupConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePostgresBackupConfig",
		Method:             "POST",
		PathPattern:        "/v1/database/postgres/backup-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePostgresBackupConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePostgresBackupConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePostgresBackupConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
