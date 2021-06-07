// Code generated by go-swagger; DO NOT EDIT.

package gateway

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

// NewClientGetParams creates a new ClientGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClientGetParams() *ClientGetParams {
	return &ClientGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClientGetParamsWithTimeout creates a new ClientGetParams object
// with the ability to set a timeout on a request.
func NewClientGetParamsWithTimeout(timeout time.Duration) *ClientGetParams {
	return &ClientGetParams{
		timeout: timeout,
	}
}

// NewClientGetParamsWithContext creates a new ClientGetParams object
// with the ability to set a context for a request.
func NewClientGetParamsWithContext(ctx context.Context) *ClientGetParams {
	return &ClientGetParams{
		Context: ctx,
	}
}

// NewClientGetParamsWithHTTPClient creates a new ClientGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClientGetParamsWithHTTPClient(client *http.Client) *ClientGetParams {
	return &ClientGetParams{
		HTTPClient: client,
	}
}

/* ClientGetParams contains all the parameters to send to the API endpoint
   for the client get operation.

   Typically these are written to a http.Request.
*/
type ClientGetParams struct {

	/* Name.

	   name of the gateway server
	*/
	Name string

	/* Projectuid.

	   project UUID
	*/
	Projectuid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the client get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClientGetParams) WithDefaults() *ClientGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the client get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClientGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the client get params
func (o *ClientGetParams) WithTimeout(timeout time.Duration) *ClientGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the client get params
func (o *ClientGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the client get params
func (o *ClientGetParams) WithContext(ctx context.Context) *ClientGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the client get params
func (o *ClientGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the client get params
func (o *ClientGetParams) WithHTTPClient(client *http.Client) *ClientGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the client get params
func (o *ClientGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the client get params
func (o *ClientGetParams) WithName(name string) *ClientGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the client get params
func (o *ClientGetParams) SetName(name string) {
	o.Name = name
}

// WithProjectuid adds the projectuid to the client get params
func (o *ClientGetParams) WithProjectuid(projectuid string) *ClientGetParams {
	o.SetProjectuid(projectuid)
	return o
}

// SetProjectuid adds the projectuid to the client get params
func (o *ClientGetParams) SetProjectuid(projectuid string) {
	o.Projectuid = projectuid
}

// WriteToRequest writes these params to a swagger request
func (o *ClientGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param projectuid
	if err := r.SetPathParam("projectuid", o.Projectuid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
