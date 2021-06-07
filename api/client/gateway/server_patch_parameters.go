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

	"github.com/fi-ts/cloud-go/api/models"
)

// NewServerPatchParams creates a new ServerPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerPatchParams() *ServerPatchParams {
	return &ServerPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerPatchParamsWithTimeout creates a new ServerPatchParams object
// with the ability to set a timeout on a request.
func NewServerPatchParamsWithTimeout(timeout time.Duration) *ServerPatchParams {
	return &ServerPatchParams{
		timeout: timeout,
	}
}

// NewServerPatchParamsWithContext creates a new ServerPatchParams object
// with the ability to set a context for a request.
func NewServerPatchParamsWithContext(ctx context.Context) *ServerPatchParams {
	return &ServerPatchParams{
		Context: ctx,
	}
}

// NewServerPatchParamsWithHTTPClient creates a new ServerPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerPatchParamsWithHTTPClient(client *http.Client) *ServerPatchParams {
	return &ServerPatchParams{
		HTTPClient: client,
	}
}

/* ServerPatchParams contains all the parameters to send to the API endpoint
   for the server patch operation.

   Typically these are written to a http.Request.
*/
type ServerPatchParams struct {

	// Body.
	Body *models.V1GatewayServerPatchRequest

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

// WithDefaults hydrates default values in the server patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerPatchParams) WithDefaults() *ServerPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the server patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the server patch params
func (o *ServerPatchParams) WithTimeout(timeout time.Duration) *ServerPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server patch params
func (o *ServerPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server patch params
func (o *ServerPatchParams) WithContext(ctx context.Context) *ServerPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server patch params
func (o *ServerPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server patch params
func (o *ServerPatchParams) WithHTTPClient(client *http.Client) *ServerPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server patch params
func (o *ServerPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the server patch params
func (o *ServerPatchParams) WithBody(body *models.V1GatewayServerPatchRequest) *ServerPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the server patch params
func (o *ServerPatchParams) SetBody(body *models.V1GatewayServerPatchRequest) {
	o.Body = body
}

// WithName adds the name to the server patch params
func (o *ServerPatchParams) WithName(name string) *ServerPatchParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the server patch params
func (o *ServerPatchParams) SetName(name string) {
	o.Name = name
}

// WithProjectuid adds the projectuid to the server patch params
func (o *ServerPatchParams) WithProjectuid(projectuid string) *ServerPatchParams {
	o.SetProjectuid(projectuid)
	return o
}

// SetProjectuid adds the projectuid to the server patch params
func (o *ServerPatchParams) SetProjectuid(projectuid string) {
	o.Projectuid = projectuid
}

// WriteToRequest writes these params to a swagger request
func (o *ServerPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
