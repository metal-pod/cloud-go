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

// V1VPN v1 v p n
//
// swagger:model v1.VPN
type V1VPN struct {

	// address
	// Required: true
	Address *string `json:"Address"`

	// auth key
	// Required: true
	AuthKey *string `json:"AuthKey"`
}

// Validate validates this v1 v p n
func (m *V1VPN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VPN) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("Address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *V1VPN) validateAuthKey(formats strfmt.Registry) error {

	if err := validate.Required("AuthKey", "body", m.AuthKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 v p n based on context it is used
func (m *V1VPN) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VPN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VPN) UnmarshalBinary(b []byte) error {
	var res V1VPN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
