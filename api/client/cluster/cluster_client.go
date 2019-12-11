// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCluster creates a cluster if the given ID already exists a conflict is returned
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCluster",
		Method:             "PUT",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterCreated), nil

}

/*
DeleteCluster deletes an cluster and returns the deleted entity
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/v1/cluster/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterOK), nil

}

/*
FindCluster gets cluster by id
*/
func (a *Client) FindCluster(params *FindClusterParams, authInfo runtime.ClientAuthInfoWriter) (*FindClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findCluster",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindClusterOK), nil

}

/*
FindClusters finds clusters by multiple criteria
*/
func (a *Client) FindClusters(params *FindClustersParams, authInfo runtime.ClientAuthInfoWriter) (*FindClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findClusters",
		Method:             "POST",
		PathPattern:        "/v1/cluster/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindClustersOK), nil

}

/*
GetClusterCredentials gets all the kubeconfig for the cluster
*/
func (a *Client) GetClusterCredentials(params *GetClusterCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterCredentials",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterCredentialsOK), nil

}

/*
GetSSHKeyPair gets all the ssh keypairs of the cluster
*/
func (a *Client) GetSSHKeyPair(params *GetSSHKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*GetSSHKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSSHKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSSHKeyPair",
		Method:             "GET",
		PathPattern:        "/v1/cluster/{id}/sshkeypair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSSHKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSSHKeyPairOK), nil

}

/*
ListClusters gets all clusters
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listClusters",
		Method:             "GET",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListConstraints gets constraints for cluster create
*/
func (a *Client) ListConstraints(params *ListConstraintsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConstraintsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListConstraintsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConstraints",
		Method:             "GET",
		PathPattern:        "/v1/cluster/constraints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListConstraintsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListConstraintsOK), nil

}

/*
ReconcileCluster triggers cluster reconcilation
*/
func (a *Client) ReconcileCluster(params *ReconcileClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ReconcileClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReconcileClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "reconcileCluster",
		Method:             "POST",
		PathPattern:        "/v1/cluster/{id}/reconcile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReconcileClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReconcileClusterOK), nil

}

/*
UpdateCluster updates a cluster if the cluster was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateCluster(params *UpdateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCluster",
		Method:             "POST",
		PathPattern:        "/v1/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
