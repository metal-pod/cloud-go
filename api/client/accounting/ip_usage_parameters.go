// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

	models "github.com/fi-ts/cloud-go/api/models"
)

// NewIPUsageParams creates a new IPUsageParams object
// with the default values initialized.
func NewIPUsageParams() *IPUsageParams {
	var ()
	return &IPUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPUsageParamsWithTimeout creates a new IPUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPUsageParamsWithTimeout(timeout time.Duration) *IPUsageParams {
	var ()
	return &IPUsageParams{

		timeout: timeout,
	}
}

// NewIPUsageParamsWithContext creates a new IPUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPUsageParamsWithContext(ctx context.Context) *IPUsageParams {
	var ()
	return &IPUsageParams{

		Context: ctx,
	}
}

// NewIPUsageParamsWithHTTPClient creates a new IPUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPUsageParamsWithHTTPClient(client *http.Client) *IPUsageParams {
	var ()
	return &IPUsageParams{
		HTTPClient: client,
	}
}

/*IPUsageParams contains all the parameters to send to the API endpoint
for the ip usage operation typically these are written to a http.Request
*/
type IPUsageParams struct {

	/*Body*/
	Body *models.V1IPUsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ip usage params
func (o *IPUsageParams) WithTimeout(timeout time.Duration) *IPUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip usage params
func (o *IPUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip usage params
func (o *IPUsageParams) WithContext(ctx context.Context) *IPUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip usage params
func (o *IPUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip usage params
func (o *IPUsageParams) WithHTTPClient(client *http.Client) *IPUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip usage params
func (o *IPUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ip usage params
func (o *IPUsageParams) WithBody(body *models.V1IPUsageRequest) *IPUsageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ip usage params
func (o *IPUsageParams) SetBody(body *models.V1IPUsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IPUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
