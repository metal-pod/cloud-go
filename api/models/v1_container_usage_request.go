// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ContainerUsageRequest v1 container usage request
//
// swagger:model v1.ContainerUsageRequest
type V1ContainerUsageRequest struct {

	// the cluster id to account for
	Clusterid string `json:"clusterid,omitempty"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// the namespace name to account for
	Namespace string `json:"namespace,omitempty"`

	// the project id to account for
	Projectid string `json:"projectid,omitempty"`

	// the tenant to get the container usage for (defaults to all tenants)
	Tenant string `json:"tenant,omitempty"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`
}

// Validate validates this v1 container usage request
func (m *V1ContainerUsageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerUsageRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ContainerUsageRequest) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerUsageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerUsageRequest) UnmarshalBinary(b []byte) error {
	var res V1ContainerUsageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
