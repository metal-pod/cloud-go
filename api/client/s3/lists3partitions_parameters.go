// Code generated by go-swagger; DO NOT EDIT.

package s3

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

// NewLists3partitionsParams creates a new Lists3partitionsParams object
// with the default values initialized.
func NewLists3partitionsParams() *Lists3partitionsParams {

	return &Lists3partitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLists3partitionsParamsWithTimeout creates a new Lists3partitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLists3partitionsParamsWithTimeout(timeout time.Duration) *Lists3partitionsParams {

	return &Lists3partitionsParams{

		timeout: timeout,
	}
}

// NewLists3partitionsParamsWithContext creates a new Lists3partitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLists3partitionsParamsWithContext(ctx context.Context) *Lists3partitionsParams {

	return &Lists3partitionsParams{

		Context: ctx,
	}
}

// NewLists3partitionsParamsWithHTTPClient creates a new Lists3partitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLists3partitionsParamsWithHTTPClient(client *http.Client) *Lists3partitionsParams {

	return &Lists3partitionsParams{
		HTTPClient: client,
	}
}

/*Lists3partitionsParams contains all the parameters to send to the API endpoint
for the lists3partitions operation typically these are written to a http.Request
*/
type Lists3partitionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lists3partitions params
func (o *Lists3partitionsParams) WithTimeout(timeout time.Duration) *Lists3partitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lists3partitions params
func (o *Lists3partitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lists3partitions params
func (o *Lists3partitionsParams) WithContext(ctx context.Context) *Lists3partitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lists3partitions params
func (o *Lists3partitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lists3partitions params
func (o *Lists3partitionsParams) WithHTTPClient(client *http.Client) *Lists3partitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lists3partitions params
func (o *Lists3partitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *Lists3partitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
