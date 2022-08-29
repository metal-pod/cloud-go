// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1PostgresUsageRequest v1 postgres usage request
//
// swagger:model v1.PostgresUsageRequest
type V1PostgresUsageRequest struct {

	// accounting annotations present on the last accounting report of this postgres
	// Required: true
	Annotations []string `json:"annotations" yaml:"annotations"`

	// the cluster id to account for
	Clusterid string `json:"clusterid,omitempty" yaml:"clusterid,omitempty"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from" yaml:"from"`

	// the project id to account for
	Projectid string `json:"projectid,omitempty" yaml:"projectid,omitempty"`

	// the tenant to get the container usage for (defaults to all tenants)
	Tenant string `json:"tenant,omitempty" yaml:"tenant,omitempty"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty" yaml:"to,omitempty"`

	// the uuid of this postgres
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Validate validates this v1 postgres usage request
func (m *V1PostgresUsageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

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

func (m *V1PostgresUsageRequest) validateAnnotations(formats strfmt.Registry) error {

	if err := validate.Required("annotations", "body", m.Annotations); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresUsageRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresUsageRequest) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 postgres usage request based on context it is used
func (m *V1PostgresUsageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresUsageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresUsageRequest) UnmarshalBinary(b []byte) error {
	var res V1PostgresUsageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
