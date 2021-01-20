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

// ModelsV1MachineBIOS models v1 machine b i o s
//
// swagger:model models.V1MachineBIOS
type ModelsV1MachineBIOS struct {

	// date
	// Required: true
	Date *string `json:"date"`

	// vendor
	// Required: true
	Vendor *string `json:"vendor"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this models v1 machine b i o s
func (m *ModelsV1MachineBIOS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineBIOS) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineBIOS) validateVendor(formats strfmt.Registry) error {

	if err := validate.Required("vendor", "body", m.Vendor); err != nil {
		return err
	}

	return nil
}

func (m *ModelsV1MachineBIOS) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 machine b i o s based on context it is used
func (m *ModelsV1MachineBIOS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineBIOS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineBIOS) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineBIOS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
