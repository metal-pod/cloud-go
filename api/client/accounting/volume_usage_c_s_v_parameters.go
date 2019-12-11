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

	models "github.com/metal-pod/cloud-go/api/models"
)

// NewVolumeUsageCSVParams creates a new VolumeUsageCSVParams object
// with the default values initialized.
func NewVolumeUsageCSVParams() *VolumeUsageCSVParams {
	var ()
	return &VolumeUsageCSVParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeUsageCSVParamsWithTimeout creates a new VolumeUsageCSVParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVolumeUsageCSVParamsWithTimeout(timeout time.Duration) *VolumeUsageCSVParams {
	var ()
	return &VolumeUsageCSVParams{

		timeout: timeout,
	}
}

// NewVolumeUsageCSVParamsWithContext creates a new VolumeUsageCSVParams object
// with the default values initialized, and the ability to set a context for a request
func NewVolumeUsageCSVParamsWithContext(ctx context.Context) *VolumeUsageCSVParams {
	var ()
	return &VolumeUsageCSVParams{

		Context: ctx,
	}
}

// NewVolumeUsageCSVParamsWithHTTPClient creates a new VolumeUsageCSVParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVolumeUsageCSVParamsWithHTTPClient(client *http.Client) *VolumeUsageCSVParams {
	var ()
	return &VolumeUsageCSVParams{
		HTTPClient: client,
	}
}

/*VolumeUsageCSVParams contains all the parameters to send to the API endpoint
for the volume usage c s v operation typically these are written to a http.Request
*/
type VolumeUsageCSVParams struct {

	/*Body*/
	Body *models.V1VolumeUsageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the volume usage c s v params
func (o *VolumeUsageCSVParams) WithTimeout(timeout time.Duration) *VolumeUsageCSVParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume usage c s v params
func (o *VolumeUsageCSVParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume usage c s v params
func (o *VolumeUsageCSVParams) WithContext(ctx context.Context) *VolumeUsageCSVParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume usage c s v params
func (o *VolumeUsageCSVParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume usage c s v params
func (o *VolumeUsageCSVParams) WithHTTPClient(client *http.Client) *VolumeUsageCSVParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume usage c s v params
func (o *VolumeUsageCSVParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the volume usage c s v params
func (o *VolumeUsageCSVParams) WithBody(body *models.V1VolumeUsageRequest) *VolumeUsageCSVParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the volume usage c s v params
func (o *VolumeUsageCSVParams) SetBody(body *models.V1VolumeUsageRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeUsageCSVParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
