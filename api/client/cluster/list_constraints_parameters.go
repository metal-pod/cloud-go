// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListConstraintsParams creates a new ListConstraintsParams object
// with the default values initialized.
func NewListConstraintsParams() *ListConstraintsParams {

	return &ListConstraintsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListConstraintsParamsWithTimeout creates a new ListConstraintsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConstraintsParamsWithTimeout(timeout time.Duration) *ListConstraintsParams {

	return &ListConstraintsParams{

		timeout: timeout,
	}
}

// NewListConstraintsParamsWithContext creates a new ListConstraintsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConstraintsParamsWithContext(ctx context.Context) *ListConstraintsParams {

	return &ListConstraintsParams{

		Context: ctx,
	}
}

// NewListConstraintsParamsWithHTTPClient creates a new ListConstraintsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConstraintsParamsWithHTTPClient(client *http.Client) *ListConstraintsParams {

	return &ListConstraintsParams{
		HTTPClient: client,
	}
}

/*ListConstraintsParams contains all the parameters to send to the API endpoint
for the list constraints operation typically these are written to a http.Request
*/
type ListConstraintsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list constraints params
func (o *ListConstraintsParams) WithTimeout(timeout time.Duration) *ListConstraintsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list constraints params
func (o *ListConstraintsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list constraints params
func (o *ListConstraintsParams) WithContext(ctx context.Context) *ListConstraintsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list constraints params
func (o *ListConstraintsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list constraints params
func (o *ListConstraintsParams) WithHTTPClient(client *http.Client) *ListConstraintsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list constraints params
func (o *ListConstraintsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListConstraintsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
