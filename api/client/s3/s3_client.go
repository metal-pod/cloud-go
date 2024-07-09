// Code generated by go-swagger; DO NOT EDIT.

package s3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new s3 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new s3 API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new s3 API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for s3 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Creates3(params *Creates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Creates3OK, error)

	Deletes3(params *Deletes3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Deletes3OK, error)

	Gets3(params *Gets3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Gets3OK, error)

	Lists3(params *Lists3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Lists3OK, error)

	Lists3partitions(params *Lists3partitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Lists3partitionsOK, error)

	Updates3(params *Updates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Updates3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Creates3 creates an s3 user if the given name for this tenant already exists a conflict is returned
*/
func (a *Client) Creates3(params *Creates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Creates3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreates3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "creates3",
		Method:             "PUT",
		PathPattern:        "/v1/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Creates3Reader{formats: a.formats},
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
	success, ok := result.(*Creates3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Creates3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Deletes3 deletes an s3 user
*/
func (a *Client) Deletes3(params *Deletes3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Deletes3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletes3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletes3",
		Method:             "DELETE",
		PathPattern:        "/v1/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Deletes3Reader{formats: a.formats},
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
	success, ok := result.(*Deletes3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Deletes3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Gets3 gets s3 user
*/
func (a *Client) Gets3(params *Gets3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Gets3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGets3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "gets3",
		Method:             "GET",
		PathPattern:        "/v1/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Gets3Reader{formats: a.formats},
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
	success, ok := result.(*Gets3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Gets3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Lists3 lists s3 users
*/
func (a *Client) Lists3(params *Lists3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Lists3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLists3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "lists3",
		Method:             "GET",
		PathPattern:        "/v1/s3/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Lists3Reader{formats: a.formats},
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
	success, ok := result.(*Lists3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Lists3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Lists3partitions lists s3 partitions
*/
func (a *Client) Lists3partitions(params *Lists3partitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Lists3partitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLists3partitionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "lists3partitions",
		Method:             "GET",
		PathPattern:        "/v1/s3/partitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Lists3partitionsReader{formats: a.formats},
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
	success, ok := result.(*Lists3partitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Lists3partitionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Updates3 updates an s3 user
*/
func (a *Client) Updates3(params *Updates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Updates3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdates3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updates3",
		Method:             "POST",
		PathPattern:        "/v1/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Updates3Reader{formats: a.formats},
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
	success, ok := result.(*Updates3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*Updates3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
