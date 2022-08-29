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

// V1MaintenanceTimeWindow v1 maintenance time window
//
// swagger:model v1.MaintenanceTimeWindow
type V1MaintenanceTimeWindow struct {

	// begin
	// Required: true
	Begin *string `json:"Begin" yaml:"Begin"`

	// end
	// Required: true
	End *string `json:"End" yaml:"End"`
}

// Validate validates this v1 maintenance time window
func (m *V1MaintenanceTimeWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MaintenanceTimeWindow) validateBegin(formats strfmt.Registry) error {

	if err := validate.Required("Begin", "body", m.Begin); err != nil {
		return err
	}

	return nil
}

func (m *V1MaintenanceTimeWindow) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("End", "body", m.End); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 maintenance time window based on context it is used
func (m *V1MaintenanceTimeWindow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1MaintenanceTimeWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MaintenanceTimeWindow) UnmarshalBinary(b []byte) error {
	var res V1MaintenanceTimeWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
