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

// V1NetworkAccessRestrictions v1 network access restrictions
//
// swagger:model v1.NetworkAccessRestrictions
type V1NetworkAccessRestrictions struct {

	// the list of networks which are allowed to configure if networkAccessTypeForbidden is specified
	// Required: true
	AllowedNetworks []string `json:"allowed_networks"`

	// list of registries which are configured to pull only strictly required container images
	// Required: true
	MaskedRegistries []string `json:"masked_registries"`
}

// Validate validates this v1 network access restrictions
func (m *V1NetworkAccessRestrictions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaskedRegistries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkAccessRestrictions) validateAllowedNetworks(formats strfmt.Registry) error {

	if err := validate.Required("allowed_networks", "body", m.AllowedNetworks); err != nil {
		return err
	}

	return nil
}

func (m *V1NetworkAccessRestrictions) validateMaskedRegistries(formats strfmt.Registry) error {

	if err := validate.Required("masked_registries", "body", m.MaskedRegistries); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 network access restrictions based on context it is used
func (m *V1NetworkAccessRestrictions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkAccessRestrictions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkAccessRestrictions) UnmarshalBinary(b []byte) error {
	var res V1NetworkAccessRestrictions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
