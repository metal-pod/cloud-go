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

	"github.com/fi-ts/cloud-go/api/models"
)

// NewListMachineReservationsParams creates a new ListMachineReservationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMachineReservationsParams() *ListMachineReservationsParams {
	return &ListMachineReservationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMachineReservationsParamsWithTimeout creates a new ListMachineReservationsParams object
// with the ability to set a timeout on a request.
func NewListMachineReservationsParamsWithTimeout(timeout time.Duration) *ListMachineReservationsParams {
	return &ListMachineReservationsParams{
		timeout: timeout,
	}
}

// NewListMachineReservationsParamsWithContext creates a new ListMachineReservationsParams object
// with the ability to set a context for a request.
func NewListMachineReservationsParamsWithContext(ctx context.Context) *ListMachineReservationsParams {
	return &ListMachineReservationsParams{
		Context: ctx,
	}
}

// NewListMachineReservationsParamsWithHTTPClient creates a new ListMachineReservationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMachineReservationsParamsWithHTTPClient(client *http.Client) *ListMachineReservationsParams {
	return &ListMachineReservationsParams{
		HTTPClient: client,
	}
}

/*
ListMachineReservationsParams contains all the parameters to send to the API endpoint

	for the list machine reservations operation.

	Typically these are written to a http.Request.
*/
type ListMachineReservationsParams struct {

	// Body.
	Body *models.V1MachineReservationFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list machine reservations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineReservationsParams) WithDefaults() *ListMachineReservationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list machine reservations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineReservationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list machine reservations params
func (o *ListMachineReservationsParams) WithTimeout(timeout time.Duration) *ListMachineReservationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list machine reservations params
func (o *ListMachineReservationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list machine reservations params
func (o *ListMachineReservationsParams) WithContext(ctx context.Context) *ListMachineReservationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list machine reservations params
func (o *ListMachineReservationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list machine reservations params
func (o *ListMachineReservationsParams) WithHTTPClient(client *http.Client) *ListMachineReservationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list machine reservations params
func (o *ListMachineReservationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list machine reservations params
func (o *ListMachineReservationsParams) WithBody(body *models.V1MachineReservationFindRequest) *ListMachineReservationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list machine reservations params
func (o *ListMachineReservationsParams) SetBody(body *models.V1MachineReservationFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListMachineReservationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
