// Code generated by go-swagger; DO NOT EDIT.

package gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gateway API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gateway API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteClient(params *DeleteClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClientOK, error)

	DeleteServer(params *DeleteServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServerOK, error)

	GetClient(params *GetClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientOK, error)

	GetServer(params *GetServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerOK, error)

	ListServers(params *ListServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServersOK, error)

	PatchClient(params *PatchClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchClientOK, error)

	PatchServer(params *PatchServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServerOK, error)

	PostClient(params *PostClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostClientCreated, error)

	PostServer(params *PostServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostServerCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteClient delete client API
*/
func (a *Client) DeleteClient(params *DeleteClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteClient",
		Method:             "DELETE",
		PathPattern:        "/gateway/v1/projects/{projectuid}/clients/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClientReader{formats: a.formats},
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
	success, ok := result.(*DeleteClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteServer delete server API
*/
func (a *Client) DeleteServer(params *DeleteServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteServer",
		Method:             "DELETE",
		PathPattern:        "/gateway/v1/projects/{projectuid}/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServerReader{formats: a.formats},
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
	success, ok := result.(*DeleteServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetClient get client API
*/
func (a *Client) GetClient(params *GetClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClient",
		Method:             "GET",
		PathPattern:        "/gateway/v1/projects/{projectuid}/clients/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClientReader{formats: a.formats},
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
	success, ok := result.(*GetClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetServer get server API
*/
func (a *Client) GetServer(params *GetServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getServer",
		Method:             "GET",
		PathPattern:        "/gateway/v1/projects/{projectuid}/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServerReader{formats: a.formats},
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
	success, ok := result.(*GetServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListServers list servers API
*/
func (a *Client) ListServers(params *ListServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listServers",
		Method:             "GET",
		PathPattern:        "/gateway/v1/projects/{projectuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServersReader{formats: a.formats},
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
	success, ok := result.(*ListServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchClient patch client API
*/
func (a *Client) PatchClient(params *PatchClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchClient",
		Method:             "PATCH",
		PathPattern:        "/gateway/v1/projects/{projectuid}/clients/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchClientReader{formats: a.formats},
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
	success, ok := result.(*PatchClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchServer patch server API
*/
func (a *Client) PatchServer(params *PatchServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchServer",
		Method:             "PATCH",
		PathPattern:        "/gateway/v1/projects/{projectuid}/servers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchServerReader{formats: a.formats},
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
	success, ok := result.(*PatchServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostClient post client API
*/
func (a *Client) PostClient(params *PostClientParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostClientCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postClient",
		Method:             "POST",
		PathPattern:        "/gateway/v1/projects/{projectuid}/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClientReader{formats: a.formats},
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
	success, ok := result.(*PostClientCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostServer post server API
*/
func (a *Client) PostServer(params *PostServerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostServerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postServer",
		Method:             "POST",
		PathPattern:        "/gateway/v1/projects/{projectuid}/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostServerReader{formats: a.formats},
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
	success, ok := result.(*PostServerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
