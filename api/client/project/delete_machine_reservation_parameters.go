// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewDeleteMachineReservationParams creates a new DeleteMachineReservationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMachineReservationParams() *DeleteMachineReservationParams {
	return &DeleteMachineReservationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMachineReservationParamsWithTimeout creates a new DeleteMachineReservationParams object
// with the ability to set a timeout on a request.
func NewDeleteMachineReservationParamsWithTimeout(timeout time.Duration) *DeleteMachineReservationParams {
	return &DeleteMachineReservationParams{
		timeout: timeout,
	}
}

// NewDeleteMachineReservationParamsWithContext creates a new DeleteMachineReservationParams object
// with the ability to set a context for a request.
func NewDeleteMachineReservationParamsWithContext(ctx context.Context) *DeleteMachineReservationParams {
	return &DeleteMachineReservationParams{
		Context: ctx,
	}
}

// NewDeleteMachineReservationParamsWithHTTPClient creates a new DeleteMachineReservationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMachineReservationParamsWithHTTPClient(client *http.Client) *DeleteMachineReservationParams {
	return &DeleteMachineReservationParams{
		HTTPClient: client,
	}
}

/*
DeleteMachineReservationParams contains all the parameters to send to the API endpoint

	for the delete machine reservation operation.

	Typically these are written to a http.Request.
*/
type DeleteMachineReservationParams struct {

	/* Project.

	   identifier of the project
	*/
	Project string

	/* Size.

	   identifier of the size
	*/
	Size string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete machine reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMachineReservationParams) WithDefaults() *DeleteMachineReservationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete machine reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMachineReservationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete machine reservation params
func (o *DeleteMachineReservationParams) WithTimeout(timeout time.Duration) *DeleteMachineReservationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete machine reservation params
func (o *DeleteMachineReservationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete machine reservation params
func (o *DeleteMachineReservationParams) WithContext(ctx context.Context) *DeleteMachineReservationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete machine reservation params
func (o *DeleteMachineReservationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete machine reservation params
func (o *DeleteMachineReservationParams) WithHTTPClient(client *http.Client) *DeleteMachineReservationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete machine reservation params
func (o *DeleteMachineReservationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProject adds the project to the delete machine reservation params
func (o *DeleteMachineReservationParams) WithProject(project string) *DeleteMachineReservationParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete machine reservation params
func (o *DeleteMachineReservationParams) SetProject(project string) {
	o.Project = project
}

// WithSize adds the size to the delete machine reservation params
func (o *DeleteMachineReservationParams) WithSize(size string) *DeleteMachineReservationParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the delete machine reservation params
func (o *DeleteMachineReservationParams) SetSize(size string) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMachineReservationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param size
	if err := r.SetPathParam("size", o.Size); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
