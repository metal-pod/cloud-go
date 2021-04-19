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

// NewDeletePostgresBackupConfigParams creates a new DeletePostgresBackupConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePostgresBackupConfigParams() *DeletePostgresBackupConfigParams {
	return &DeletePostgresBackupConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePostgresBackupConfigParamsWithTimeout creates a new DeletePostgresBackupConfigParams object
// with the ability to set a timeout on a request.
func NewDeletePostgresBackupConfigParamsWithTimeout(timeout time.Duration) *DeletePostgresBackupConfigParams {
	return &DeletePostgresBackupConfigParams{
		timeout: timeout,
	}
}

// NewDeletePostgresBackupConfigParamsWithContext creates a new DeletePostgresBackupConfigParams object
// with the ability to set a context for a request.
func NewDeletePostgresBackupConfigParamsWithContext(ctx context.Context) *DeletePostgresBackupConfigParams {
	return &DeletePostgresBackupConfigParams{
		Context: ctx,
	}
}

// NewDeletePostgresBackupConfigParamsWithHTTPClient creates a new DeletePostgresBackupConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePostgresBackupConfigParamsWithHTTPClient(client *http.Client) *DeletePostgresBackupConfigParams {
	return &DeletePostgresBackupConfigParams{
		HTTPClient: client,
	}
}

/* DeletePostgresBackupConfigParams contains all the parameters to send to the API endpoint
   for the delete postgres backup config operation.

   Typically these are written to a http.Request.
*/
type DeletePostgresBackupConfigParams struct {

	/* ID.

	   identifier of the projectid
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete postgres backup config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePostgresBackupConfigParams) WithDefaults() *DeletePostgresBackupConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete postgres backup config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePostgresBackupConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) WithTimeout(timeout time.Duration) *DeletePostgresBackupConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) WithContext(ctx context.Context) *DeletePostgresBackupConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) WithHTTPClient(client *http.Client) *DeletePostgresBackupConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) WithID(id string) *DeletePostgresBackupConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete postgres backup config params
func (o *DeletePostgresBackupConfigParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePostgresBackupConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
