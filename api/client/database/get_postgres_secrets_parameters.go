// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPostgresSecretsParams creates a new GetPostgresSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPostgresSecretsParams() *GetPostgresSecretsParams {
	return &GetPostgresSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPostgresSecretsParamsWithTimeout creates a new GetPostgresSecretsParams object
// with the ability to set a timeout on a request.
func NewGetPostgresSecretsParamsWithTimeout(timeout time.Duration) *GetPostgresSecretsParams {
	return &GetPostgresSecretsParams{
		timeout: timeout,
	}
}

// NewGetPostgresSecretsParamsWithContext creates a new GetPostgresSecretsParams object
// with the ability to set a context for a request.
func NewGetPostgresSecretsParamsWithContext(ctx context.Context) *GetPostgresSecretsParams {
	return &GetPostgresSecretsParams{
		Context: ctx,
	}
}

// NewGetPostgresSecretsParamsWithHTTPClient creates a new GetPostgresSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPostgresSecretsParamsWithHTTPClient(client *http.Client) *GetPostgresSecretsParams {
	return &GetPostgresSecretsParams{
		HTTPClient: client,
	}
}

/* GetPostgresSecretsParams contains all the parameters to send to the API endpoint
   for the get postgres secrets operation.

   Typically these are written to a http.Request.
*/
type GetPostgresSecretsParams struct {

	/* ID.

	   identifier of the postgres
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get postgres secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPostgresSecretsParams) WithDefaults() *GetPostgresSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get postgres secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPostgresSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get postgres secrets params
func (o *GetPostgresSecretsParams) WithTimeout(timeout time.Duration) *GetPostgresSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get postgres secrets params
func (o *GetPostgresSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get postgres secrets params
func (o *GetPostgresSecretsParams) WithContext(ctx context.Context) *GetPostgresSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get postgres secrets params
func (o *GetPostgresSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get postgres secrets params
func (o *GetPostgresSecretsParams) WithHTTPClient(client *http.Client) *GetPostgresSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get postgres secrets params
func (o *GetPostgresSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get postgres secrets params
func (o *GetPostgresSecretsParams) WithID(id string) *GetPostgresSecretsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get postgres secrets params
func (o *GetPostgresSecretsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPostgresSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
