// Code generated by go-swagger; DO NOT EDIT.

package s3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new s3 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s3 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Creates3 creates an s3 user if the given name for this tenant already exists a conflict is returned
*/
func (a *Client) Creates3(params *Creates3Params, authInfo runtime.ClientAuthInfoWriter) (*Creates3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreates3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
	if err != nil {
		return nil, err
	}
	return result.(*Creates3OK), nil

}

/*
Deletes3 deletes an s3 user
*/
func (a *Client) Deletes3(params *Deletes3Params, authInfo runtime.ClientAuthInfoWriter) (*Deletes3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletes3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletes3",
		Method:             "DELETE",
		PathPattern:        "/v1/s3/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Deletes3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Deletes3OK), nil

}

/*
Gets3 gets s3 user
*/
func (a *Client) Gets3(params *Gets3Params, authInfo runtime.ClientAuthInfoWriter) (*Gets3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGets3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "gets3",
		Method:             "GET",
		PathPattern:        "/v1/s3/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Gets3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Gets3OK), nil

}

/*
Lists3 lists s3 users
*/
func (a *Client) Lists3(params *Lists3Params, authInfo runtime.ClientAuthInfoWriter) (*Lists3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLists3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lists3",
		Method:             "GET",
		PathPattern:        "/v1/s3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Lists3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Lists3OK), nil

}

/*
Lists3partitions lists s3 partitions
*/
func (a *Client) Lists3partitions(params *Lists3partitionsParams, authInfo runtime.ClientAuthInfoWriter) (*Lists3partitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLists3partitionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
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
	})
	if err != nil {
		return nil, err
	}
	return result.(*Lists3partitionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
